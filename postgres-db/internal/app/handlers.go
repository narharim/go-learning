package app

import (
	"encoding/json"
	"net/http"
)

func (a *App) Ping(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Thank you for checking up on me :)"))
}

type AuthorReq struct {
	Name string `json:"name"`
}

func (a *App) createAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var req AuthorReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		HandleAppError(w, err)
	}
}

func (a *App) listAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Thank you for checking up on me :)"))
}
