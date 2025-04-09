package main

import (
	"fmt"
	userProto "github.com/Errera11/authorization/internal/protogen/user"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Errorf("Error loading .env file")
	}

	// register redis store
	redisConn := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	// register user service
	conn, err := grpc.NewClient(os.Getenv("USER_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	userService := userProto.NewUserServiceClient(conn)

	if err != nil {
		log.Fatalf("could not regisiter userService: %v", err)
	}

	server := NewGRPCServer(os.Getenv("SERVICE_URL"), redisConn, userService)

	server.Run()
}
