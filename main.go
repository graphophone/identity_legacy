package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"graphophone.identity/internal/config"
	"graphophone.identity/internal/services"
	"graphophone.identity/internal/services/common/user"
)

const (
	configPath = "config.local.yaml"
)

func main() {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("Error while reading config:", err)
	}

	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "config", cfg)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error while creating listener:", err)
	}

	defer func() {
		if err := lis.Close(); err != nil {
			log.Fatal("Erorr while closing tcp listener:", err)
		}
	}()

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, services.NewUserServer(ctx))

	log.Print("Serving grpc on port 8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Error while serving grpc api:", err)
	}
}
