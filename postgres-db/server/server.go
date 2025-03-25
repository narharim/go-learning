package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port string
}

type server struct {
	Config *config
}

func New(cfg *config) *server {
	return &server{cfg}
}

func NewConfig() *config {
	return &config{
		port: os.Getenv("APP_PORT"),
	}
}

func (s *server) Start() error {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", s.Config.port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Starting server on %s", server.Addr)
	return server.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	if s != nil {
		return s.Shutdown(ctx)
	}
	return nil
}
