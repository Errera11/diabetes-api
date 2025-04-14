package main

import (
	"fmt"
	"github.com/Errera1/prediction/internal/prediction/infrastructure/db"
	"github.com/Errera1/prediction/internal/prediction/infrastructure/handler"
	"github.com/Errera1/prediction/internal/prediction/middleware"
	"github.com/Errera1/prediction/internal/prediction/service"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GRPCServer struct {
	addr        string
	storageConn *pgx.Conn
	apiAddr     string
}

func NewGRPCServer(addr string, storageConn *pgx.Conn, apiAddr string) *GRPCServer {
	fmt.Println("Creating new gRPCServer")
	return &GRPCServer{addr: addr, storageConn: storageConn, apiAddr: apiAddr}
}

func (s *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	apiService := service.NewApiService(s.apiAddr)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.UnaryServerInterceptor(middleware.MyAuthFunc)))
	reflection.Register(grpcServer)

	predictionStorage := db.New(s.storageConn)
	predictionService := service.New(predictionStorage, apiService)
	handler.New(grpcServer, predictionService)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}
