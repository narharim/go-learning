package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) RegisterRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/ping", a.Ping).Methods("GET")
	return r
}

func (a *App) Ping(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Thank you for checking up on me :)"))
}
