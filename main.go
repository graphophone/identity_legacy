package main

import (
	"context"
	"fmt"
	"log"

	"graphophone.identity/internal/config"
	"graphophone.identity/internal/database/postgres"
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
		log.Fatal("Error while connecting to redis: ", err)
	}
	fmt.Println("Connected to Redis")
	if err := redisClient.Close(); err != nil {
		log.Fatal("Error while closing redis connection: ", err)
	}
	fmt.Println("Disconnected from Redis")

	fmt.Printf(
		"Postgres: %s:%s (user: %s, password: %s)\n",
		cfg.Postgres().Host,
		cfg.Postgres().Port,
		cfg.Postgres().User,
		cfg.Postgres().Password,
	)
	postgresClient, err := postgres.Connect(cfg.Postgres())
	if err != nil {
		log.Fatal("Error while connecting to postgres: ", err)
	}
	fmt.Println("Connected to Postgres")
	if err := postgresClient.Close(); err != nil {
		log.Fatal("Error while closing Postgres connection: ", err)
	}
	fmt.Println("Disconnected from Postgres")
}
