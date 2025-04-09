package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gateway "github.com/Errera11/api-gateway/internal/protogen"
)

type MyGrpcServer struct{}

func (s *MyGrpcServer) Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := gateway.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, os.Getenv("AUTHORIZATION_MICROSERVICE_URL"), opts)
	if err != nil {
		return err
	}
	err = gateway.RegisterUserServiceHandlerFromEndpoint(ctx, mux, os.Getenv("USER_MICROSERVICE_URL"), opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Println("Starting gRPC server on", os.Getenv("GATEWAY_URL"))
	return http.ListenAndServe(os.Getenv("GATEWAY_URL"), mux)
}

func NewGrpcSever() *MyGrpcServer {
	return &MyGrpcServer{}
}
