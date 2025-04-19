package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Errorf("Error loading .env file")
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil || conn == nil {
		fmt.Println(err)
		panic(fmt.Errorf("Error connecting to database: %v", err))
	}

	server := NewGRPCServer(os.Getenv("SERVICE_URL"), conn)

	server.Run()
}
