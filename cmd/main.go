package main

import (
	"context"
	"fmt"
	"log"

	"graphophone.identity/internal/config"
	"graphophone.identity/internal/database/redis"
)

func main() {
	ctx := context.Background()
	fmt.Println("Hello world")
	cfg, err := config.LoadConfig("config.local.yaml")
	if err != nil {
		log.Fatal("Error while reading config: ", err)
		return
	}
	fmt.Printf("Redis: %s (password: %s)\n", cfg.Redis().Address, cfg.Redis().Password)
	redisClient, err := redis.Connect(ctx, cfg.Redis())
	if err != nil {
		log.Fatal("Error while reading config: ", err)
		return
	}

	if err := redisClient.Close(); err != nil {
		log.Fatal("Error while closing redis connection: ", err)
		return
	}
}
