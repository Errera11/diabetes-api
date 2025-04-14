package middleware

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"os"
	"strings"

	protogen "github.com/Errera1/prediction/internal/protogen"
)

type AuthFunc func(ctx context.Context) (context.Context, error)

type ServiceAuthFuncOverride interface {
	AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error)
}

var publicRoutes = []string{
	"GetUserByEmail",
	"GetUserById",
}

func UnaryServerInterceptor(authFunc AuthFunc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		for _, route := range publicRoutes {
			if strings.Contains(info.FullMethod, route) {
				return handler(ctx, req)
			}
		}

		var newCtx context.Context
		var err error
		if overrideSrv, ok := info.Server.(ServiceAuthFuncOverride); ok {
			newCtx, err = overrideSrv.AuthFuncOverride(ctx, info.FullMethod)
		} else {
			newCtx, err = authFunc(ctx)
		}
		if err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}

func MyAuthFunc(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	token := md.Get("userid")

	if !ok || len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "unauthorized: missing or invalid token")
	}

	user, err := ReqForAuthCheck(ctx, token[0])

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "unauthorized: missing or invalid token")
	}

	ctx = context.WithValue(ctx, "user", user)

	return ctx, nil
}

func ReqForAuthCheck(ctx context.Context, token string) (*protogen.AuthResponse, error) {
	grpcConn, err := grpc.NewClient(os.Getenv("AUTH_MICROSERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return &protogen.AuthResponse{}, fmt.Errorf("could not create grpc connection: %w", err)
	}

	authService := protogen.NewAuthServiceClient(grpcConn)

	resp, err := authService.Auth(ctx, &protogen.AuthRequest{
		Token: token,
	})

	if err != nil {
		return &protogen.AuthResponse{}, fmt.Errorf("Error during authenticatiing %w", err)
	}

	return resp, nil
}
