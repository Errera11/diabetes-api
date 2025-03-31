package main

import (
	"fmt"
	"github.com/Errera11/user/internal/user/infrastructure"
	"github.com/Errera11/user/internal/user/infrastructure/handler"
	"github.com/Errera11/user/internal/user/service"
	"github.com/jackc/pgx/v5"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	addr   string
	dbConn *pgx.Conn
}

func NewGRPCServer(addr string, dbConn *pgx.Conn) *GRPCServer {
	fmt.Println("Creating new gRPCServer")
	return &GRPCServer{addr: addr, dbConn: dbConn}
}

func (s *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	userRepo := infrastructure.NewUserRepo(s.dbConn)
	userService := service.NewUserService(userRepo)
	handler.NewGrpcUserService(grpcServer, *userService)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}
