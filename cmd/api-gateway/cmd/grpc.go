package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/rs/cors"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gateway "github.com/Errera11/api-gateway/internal/protogen"
)

type Session struct {
	UserID string
}

type contextKey int

const (
	// name of session we're setting
	defaultSessionID = "sid"

	// secret key used to encrypt session
	cookieStorageKey = "asdflkjasflkjasldfkjs"

	// How long our profile session can last, in seconds, unless renewed.
	sessionLength = 60 * 60 * 4

	// our unique key used for storing the request in the context
	requestContextKey contextKey = 0
)

type MyGrpcServer struct {
	AuthClient gateway.AuthServiceClient
}

var sessionStore sessions.Store

type gatewayMiddleware struct {
}

func (middleware *gatewayMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if ctx == nil {
			ctx = context.Background()
		}
		ctx = context.WithValue(ctx, requestContextKey, r)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getRequestFromContext(ctx context.Context) *http.Request {
	return ctx.Value(requestContextKey).(*http.Request)
}

func firstMetadataWithName(md runtime.ServerMetadata, name string) string {
	values := md.HeaderMD.Get(name)
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

func getUserIDFromServerMetadata(md runtime.ServerMetadata) (string, error) {
	userIDString := firstMetadataWithName(md, "sid")
	if userIDString == "" {
		return "", nil
	}
	if userIDString != "" {
		return userIDString, nil
	}
	return "", nil
}

func getBoolFromServerMetadata(md runtime.ServerMetadata, name string, defaultValue bool) (bool, error) {
	boolString := firstMetadataWithName(md, name)
	if boolString != "" {
		value, err := strconv.ParseBool(boolString)
		if err != nil {
			return defaultValue, err
		}
		return value, nil
	}
	return defaultValue, nil
}

// look up session and pass userId in to context if it exists
func gatewayMetadataAnnotator(ctx context.Context, r *http.Request) metadata.MD {
	session, err := sessionStore.Get(r, defaultSessionID)
	if err != nil {
		// no session, or invalid session, so pass along no extra metadata
		return metadata.Pairs()
	}
	if userIDSessionValue, ok := session.Values["userId"]; ok {
		// convert back to a Session
		userIDSession := userIDSessionValue.(*Session)
		userID := userIDSession.UserID
		// set user ID from session in the gRPC metadata
		return metadata.Pairs("userId", userID)
	}
	// otherwise pass no extra metadata along
	return metadata.Pairs()
}

func gatewayResponseModifier(ctx context.Context, response http.ResponseWriter, _ proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return fmt.Errorf("Failed to extract ServerMetadata from context")
	}
	// did the gRPC method set a user ID in the metadata?
	userID, err := getUserIDFromServerMetadata(md)
	if err != nil {
		return err
	}

	if userID != "" {
		//rlog.Debugf("gRPC call set userId to %d", userID)

		// pull the request from context (set in middleware above)
		request := getRequestFromContext(ctx)

		// create or get the session
		session, err := sessionStore.New(request, defaultSessionID)

		if err != nil {
			//rlog.Error(err, "couldn't create a session")
			return err
		}
		session.Options.MaxAge = sessionLength
		session.Options.Path = "/"
		session.Options.HttpOnly = true
		session.Options.SameSite = http.SameSiteNoneMode
		session.Options.Secure = false

		// create a session for the user.  This session is converted by gorilla
		// into a session cookie
		userIDSession := &Session{
			UserID: userID,
		}

		// put the userId into session
		session.Values["userId"] = userIDSession
		// save the session, creating a cookie from it
		if err := sessionStore.Save(request, response, session); err != nil {
			//rlog.Error(err, "couldn't save the session as a cookie")
			return err
		}
	}

	// did the gRPC method called set a flag telling us to delete the session?
	deleteSession, err := getBoolFromServerMetadata(md, "sid-del", false)
	if err != nil {
		return err
	}
	if deleteSession {
		// pull the request from context (set in middleware above)
		r := getRequestFromContext(ctx)

		// as documented, to delete session, set max age to -1
		session, err := sessionStore.New(r, defaultSessionID)
		if err != nil {
			//rlog.Error(err, "couldn't create empty session")
			return err
		}
		session.Options.MaxAge = -1
		session.Options.Path = "/"
		// "save" the session with maxage = -1, clearing it
		if err := sessionStore.Save(r, response, session); err != nil {
			//rlog.Error(err, "couldn't delete session")
			return err
		}
	}

	return nil
}

func (s *MyGrpcServer) Run() error {
	// session handling
	gob.Register(&Session{})

	cookieStore := sessions.NewCookieStore([]byte(cookieStorageKey))
	cookieStore.Options = &sessions.Options{HttpOnly: true}
	sessionStore = cookieStore

	r := mux.NewRouter()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux(runtime.WithForwardResponseOption(gatewayResponseModifier), runtime.WithMetadata(gatewayMetadataAnnotator))
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := gateway.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, os.Getenv("AUTHORIZATION_MICROSERVICE_URL"), opts)
	if err != nil {
		return err
	}
	err = gateway.RegisterUserServiceHandlerFromEndpoint(ctx, mux, os.Getenv("USER_MICROSERVICE_URL"), opts)
	if err != nil {
		return err
	}
	err = gateway.RegisterPredictionServiceHandlerFromEndpoint(ctx, mux, os.Getenv("PREDICTION_MICROSERVICE_URL"), opts)
	if err != nil {
		return err
	}

	loggingHandler := handlers.CombinedLoggingHandler(os.Stderr, mux)
	handlers.AllowedOrigins([]string{"*"})
	middleware := gatewayMiddleware{}
	r.Use(middleware.Middleware)
	r.PathPrefix("/").Handler(loggingHandler)

	withCors := cors.New(cors.Options{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(r)

	gatewayAddr := os.Getenv("GATEWAY_URL")
	srv := &http.Server{
		Addr:         gatewayAddr,
		WriteTimeout: time.Second * 20,
		ReadTimeout:  time.Second * 20,
		IdleTimeout:  time.Second * 65,
		Handler:      withCors,
	}

	log.Println("Starting api-gateway http server on", gatewayAddr)

	return srv.ListenAndServe()
}

func NewGrpcSever() *MyGrpcServer {
	return &MyGrpcServer{}
}
