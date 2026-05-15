package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConnectPostgres() {

	connStr := "postgres://postgres:database123@localhost:5432/postgres"

	db, err := sql.Open("pgx", connStr)

	if err != nil {
		log.Fatal("Error while connecting:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Database not responding:", err)
	}

	DB = db

	fmt.Println("PostgreSQL Connected Successfully")
}