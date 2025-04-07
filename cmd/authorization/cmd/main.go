package main

import "github.com/redis/go-redis/v9"

func main() {
	redisConn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	server := NewGRPCServer(":4321", redisConn)

	server.Run()
}
