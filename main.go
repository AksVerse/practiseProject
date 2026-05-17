package main

import (
	"project/config"
	"project/models"
	"project/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	config.ConnectPostgres()
	config.ConnectMongo()
	config.ConnectRedis()

	config.DB.AutoMigrate(&models.Student{})

	routes.SetupRoutes(router)

	router.Run(":8080")
}