package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/narharim/go-learning/postgres-db/database"
)

func (a *App) Ping(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Thank you for checking up on me :)"))
}

type AuthorReq struct {
	Name string `json:"name"`
	Bio  string `json:"bio,omitempty"`
}

func (a *App) createAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var req AuthorReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		writeJSONResponse(w, http.StatusBadRequest, errors.New("err-invalid-request-body"))
		return
	}
	if req.Name == "" {
		writeJSONResponse(w, http.StatusBadRequest, errors.New("err-name-is-required"))
		return
	}

	tx, err := a.db.BeginTx(r.Context(), nil)
	if err != nil {
		log.Println(err)
		writeJSONResponse(w, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
		return
	}

	q := a.dbQueries.WithTx(tx)

	defer func() {
		if p := recover(); p != nil {
			// If there's a panic, rollback the transaction
			tx.Rollback() //TODO: Need to check this
		} else if err != nil {
			tx.Rollback() // Rollback transaction on error
		}
	}()

	authorParams := database.CreateAuthorParams{
		Name: req.Name,
		Bio: sql.NullString{
			String: req.Bio,
			Valid:  req.Bio != "",
		},
	}

	_, err = q.CreateAuthor(r.Context(), authorParams)
	if err != nil {
		log.Println(err)
		writeJSONResponse(w, http.StatusInternalServerError, errors.New("err-failed-to-create-author"))
		return
	}
	writeJSONResponse(w, http.StatusCreated, "")
}

func (a *App) listAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Thank you for checking up on me :)"))
}
