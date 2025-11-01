package main

import (
	"log"

	"backend/internal/config"
	"backend/internal/server"
	"backend/internal/db"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Create and start server
	srv := server.New(cfg)

	db.Migrate()
	
	log.Printf("Starting server on port %s", cfg.Server.Port)
	if err := srv.Start(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}