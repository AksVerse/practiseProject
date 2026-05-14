package main

import (
	"log"
	"project/config"
)

func main() {
	// config.ConnectRedis()
	// config.Mong()
	// config.Post()

	err := config.ConnectRedis()

	if err != nil {
		log.Fatal(err)
	}

	println("Server Started")
}