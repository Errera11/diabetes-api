package main

import (
	"fmt"
	"github.com/Errera11/authorization/internal/authorization/infrastructure/handler"
	"github.com/Errera11/authorization/internal/authorization/service"
	userProto "github.com/Errera11/authorization/internal/protogen/user"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"github.com/Errera11/authorization/internal/authorization/infrastructure/repository"
)

type GRPCServer struct {
	addr        string
	storageConn *redis.Client
	userService userProto.UserServiceClient
}

func NewGRPCServer(addr string, storageConn *redis.Client, userService userProto.UserServiceClient) *GRPCServer {
	fmt.Println("Creating new gRPCServer")
	return &GRPCServer{addr: addr, storageConn: storageConn, userService: userService}
}

func (s *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	authorizationRepo := repository.NewAuthorizationRepo(s.storageConn)
	authorizationService := service.New(authorizationRepo, s.userService)
	handler.NewGrpcAuthorizationService(grpcServer, *authorizationService)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}
