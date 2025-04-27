package main

import (
	"fmt"
	"github.com/Errera11/user/internal/middleware"
	"github.com/Errera11/user/internal/user/infrastructure"
	"github.com/Errera11/user/internal/user/infrastructure/handler"
	"github.com/Errera11/user/internal/user/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GRPCServer struct {
	addr   string
	dbConn *pgxpool.Pool
}

func NewGRPCServer(addr string, dbConn *pgxpool.Pool) *GRPCServer {
	fmt.Println("Creating new gRPCServer")
	return &GRPCServer{addr: addr, dbConn: dbConn}
}

func (s *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.UnaryServerInterceptor(middleware.MyAuthFunc)))

	reflection.Register(grpcServer)

	userRepo, err := infrastructure.NewUserRepo(s.dbConn)

	if err != nil {
		return err
	}

	userService := service.NewUserService(userRepo)
	handler.NewGrpcUserService(grpcServer, *userService)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}
