package main

import (
	"fmt"
	"log"

	"graphophone.identity/internal/config"
)

func main() {
	fmt.Println("Hello world")
	config, err := config.LoadConfig("config.local.yaml")
	if err != nil {
		log.Fatal("Error while reading config: ", err)
	}
	fmt.Println(config.String("config-test"))
}
