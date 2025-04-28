package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Errorf("Error loading .env file")
	}

	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Errorf("Error connecting to database: %v", err)
	}

	serviceUrl := os.Getenv("PREDICTION_API_ADDR")
	server := NewGRPCServer(os.Getenv("SERVICE_URL"), conn, &serviceUrl)

	server.Run()
}
