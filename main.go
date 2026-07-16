package main

import (
	"context"
	"fmt"
	"log"

	"graphophone.identity/internal/config"
	"graphophone.identity/internal/database/postgres"
	userdb "graphophone.identity/internal/database/postgres/user"
	"graphophone.identity/internal/database/redis"
)

const (
	configPath = "config.local.yaml"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("Error while reading config: ", err)
	}

	// Redis
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

	// Postgres
	fmt.Printf(
		"Postgres: %s:%d (user: %s, password: %s)\n",
		cfg.Postgres().Host,
		cfg.Postgres().Port,
		cfg.Postgres().User,
		cfg.Postgres().Password,
	)
	pgClient, err := postgres.New(cfg.Postgres())
	if err != nil {
		log.Fatal("Error while connecting to postgres: ", err)
	}
	fmt.Println("Connected to Postgres")
	userManager, err := userdb.New(pgClient)
	if err != nil {
		log.Fatal("Error while creating user manager: ", err)
	}
	user, err := userManager.Create(ctx, &userdb.User{
		Username:     "justkinou",
		Email:        "justkinou@proton.me",
		PasswordHash: "somehash",
		FirstName:    "No",
		LastName:     "Name",
	})
	if err != nil {
		log.Fatal("Error while adding user: ", err)
	}
	fmt.Println("User created: ", user)
	if user, err := userManager.Get(ctx, user.ID); err != nil {
		log.Fatal("Error while getting user: ", err)
	} else {
		fmt.Println("User read: ", user)
	}
	user.Username = "plainkinou"
	if err := userManager.Update(ctx, user); err != nil {
		log.Fatal("Error while updating user: ", err)
	}
	fmt.Println("User updated: ", user)
	if err := userManager.UpdateIsActive(ctx, user.ID, false); err != nil {
		log.Fatal("Error while updating user is_active: ", err)
	}
	if err := userManager.UpdatePasswordHash(ctx, user.ID, "another_hash"); err != nil {
		log.Fatal("Error while updating user password_hash: ", err)
	}
	if err := userManager.UpdateAvatar(ctx, user.ID, "some_avatar_url"); err != nil {
		log.Fatal("Error while updating user avatar_url: ", err)
	}
	if user, err := userManager.Get(ctx, user.ID); err != nil {
		log.Fatal("Error while getting user: ", err)
	} else {
		fmt.Println("User after updates: ", user)
	}
	if err := pgClient.Close(); err != nil {
		log.Fatal("Error while closing Postgres connection: ", err)
	}
	fmt.Println("Disconnected from Postgres")
}
