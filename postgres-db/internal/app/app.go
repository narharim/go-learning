package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/narharim/go-learning/postgres-db/database"
)

type Config struct {
	DBConfig *database.DBConfig
	Port     string
}

type App struct {
	config     *Config
	httpServer *http.Server
	db         *database.DB
	dbQueries  *database.Queries
}

func NewApp() (*App, error) {

	cfg := &Config{database.NewConfig(), os.Getenv("APP_PORT")}

	db, err := database.NewDB(cfg.DBConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	app := &App{
		config:    cfg,
		db:        db,
		dbQueries: dbQueries,
	}

	app.httpServer = &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      app.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return app, nil
}

func (a *App) Start() error {
	log.Printf("starting server on %s", a.httpServer.Addr)

	//https://dev.to/mokiat/proper-http-shutdown-in-go-3fji
	go func() {
		if err := a.httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server error: %v", err)
		}
		log.Println("stopped serving new connections.")
	}()

	log.Println("server started successfully")
	return a.Shutdown()
}

func (a *App) Shutdown() error {

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	log.Println("shutting down server...")

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := a.httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("error during server shutdown: %v", err)
	}

	log.Println("server shutdown complete")

	return nil
}
