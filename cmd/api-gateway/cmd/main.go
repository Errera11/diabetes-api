package main

import (
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Errorf("Error loading .env file")
	}

	server := NewGrpcSever()

	server.Run()
}
