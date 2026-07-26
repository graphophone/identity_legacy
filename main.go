package main

import (
	"context"
	"fmt"
	"log"

	"graphophone.identity/internal/config"
	"graphophone.identity/internal/core/user"
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
	userDb, err := userdb.New(pgClient)
	if err != nil {
		log.Fatal("Error while creating user db: ", err)
	}
	um := user.New(userDb)
	u, err := um.Register(ctx, &user.RegisterUserData{
		Username: "justkinou",
		Email:    "justkinou@proton.me",
		Password: "myPassword",
	})
	if err != nil {
		log.Fatal("Error while registering user: ", err)
	}
	fmt.Println("User created: ", u)
	if profile, err := um.GetProfile(ctx, u.Id); err != nil {
		log.Fatal("Error while getting user: ", err)
	} else {
		fmt.Println("User profile: ", profile)
	}
	un := "No"
	ln := "Name"
	if err := um.UpdateProfile(ctx, &user.UpdateProfileData{
		Id:        u.Id,
		Username:  u.Username,
		Bio:       "My bio",
		FirstName: &un,
		LastName:  &ln,
	}); err != nil {
		log.Fatal("Error while updating user: ", err)
	}
	fmt.Println("User updated")
	if err := um.Deactive(ctx, u.Id); err != nil {
		log.Fatal("Error while deactivating user profile: ", err)
	}
	if err := um.Activate(ctx, u.Id); err != nil {
		log.Fatal("Error while activating user profile: ", err)
	}
	if err := um.UpdatePassword(ctx, u.Id, "myPassword", "myNewpassword"); err != nil {
		log.Fatal("Error while updating user password: ", err)
	}
	if err := um.UpdateAvatar(ctx, u.Id, "some_avatar_url"); err != nil {
		log.Fatal("Error while updating user avatar_url: ", err)
	}
	if user, err := um.GetProfile(ctx, u.Id); err != nil {
		log.Fatal("Error while getting user: ", err)
	} else {
		fmt.Println("User after updates: ", user)
	}
	if err := pgClient.Close(); err != nil {
		log.Fatal("Error while closing Postgres connection: ", err)
	}
	fmt.Println("Disconnected from Postgres")
}
