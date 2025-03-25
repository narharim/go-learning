package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
	"github.com/narharim/go-learning/postgres-db/database"
	"github.com/narharim/go-learning/postgres-db/server"
)

func main() {

	dbcfg := database.NewConfig()
	if err := dbcfg.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	db, err := database.NewDB(dbcfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	srvcfg := server.NewConfig()

	srv := server.New(srvcfg)

	//New to make this using context
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := srv.Start(); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	log.Println("Server started successfully")

	<-shutdown
	log.Println("Shutting down server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("Error during server shutdown: %v", err)
	}

	log.Println("Server shutdown complete")
}
