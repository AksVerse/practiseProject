package main

import (
	// "log"
	"project/config"
)

func main() {
	//config.ConnectRedis()
	config.ConnectRedis()
	config.ConnectPostgres()
	// config.Mong()
	// config.Post()

	// err := config.ConnectRedis()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// println("Server Started")
}