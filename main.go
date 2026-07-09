package main

import (
	"context"
	"fmt"
	"log"

	"graphophone.identity/internal/config"
	"graphophone.identity/internal/database/redis"
)

const (
	configPath = "config.local.yaml"
)

func main() {
	ctx := context.Background()
	fmt.Println("Hello world")
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("Error while reading config: ", err)
	}

	fmt.Printf("Redis: %s (password: %s)\n", cfg.Redis().Address, cfg.Redis().Password)
	redisClient, err := redis.Connect(ctx, cfg.Redis())
	if err != nil {
		log.Fatal("Error while reading config: ", err)
	}

	fmt.Println("Connected to Redis")

	if err := redisClient.Close(); err != nil {
		log.Fatal("Error while closing redis connection: ", err)
	}
}
