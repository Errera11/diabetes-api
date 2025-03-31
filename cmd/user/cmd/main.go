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

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Errorf("Error loading .env file")
	}
	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	//if err != nil {
	//	fmt.Errorf("Unable to connect to database: %v\n", err)
	//}

	server := NewGRPCServer(":3212", conn)

	server.Run()

	//userRepo := user.NewUserRepo(conn)

}
