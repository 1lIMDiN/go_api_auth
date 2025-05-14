package main

import (
	"log"

	"auth/pkg/database"
	"auth/pkg/server"
)

func main() {
	dbFile := "database.db"
	if err := database.Init(dbFile); err != nil {
		log.Fatalf("failed to Init database: %w", err)
	}
	defer database.DB.Close()

	port := 8080
	if err := server.Run(port); err != nil {
		log.Fatalf("failed to listen and serve: %w", err)
	}
}