package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres() {
	dsn := "postgres://postgres:database123@localhost:5432/postgres"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Postgres connection failed")
	}

	fmt.Println("Postgres connected")

	DB = database
}