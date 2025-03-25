package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/narharim/go-learning/postgres-db/internal/app"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		log.Fatalf("unable to initialize application %v", err)
		os.Exit(1)
	}

	if err := app.Start(); err != nil {
		log.Fatalf("unable to start application %v", err)
		os.Exit(1)
	}

}
