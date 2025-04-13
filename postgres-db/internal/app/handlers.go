package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	authorParams := database.CreateAuthorParams{
		Name: req.Name,
		Bio: sql.NullString{
			String: req.Bio,
			Valid:  req.Bio != "",
		},
	}

	_, err := a.dbQueries.CreateAuthor(r.Context(), authorParams)
	if err != nil {
		log.Println(err)
		writeJSONResponse(w, http.StatusInternalServerError, errors.New("err-failed-to-create-author"))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a *App) listAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	authors, err := a.dbQueries.ListAuthors(r.Context())
	if err != nil {
		log.Println(err)
		writeJSONResponse(w, http.StatusInternalServerError, errors.New("err-failed-to-list-authors"))
		return
	}
	writeJSONResponse(w, http.StatusOK, authors)
}

func (a *App) getAuthorHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, "err-invalid-id-format")
		return
	}

	authors, err := a.dbQueries.GetAuthor(r.Context(), int32(id))
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			writeJSONResponse(w, http.StatusInternalServerError, errors.New("err-author-id-not-found"))
			return
		}
		writeJSONResponse(w, http.StatusInternalServerError, errors.New("err-failed-to-get-author"))
		return
	}

	writeJSONResponse(w, http.StatusOK, authors)
}

func (a *App) updateAuthorHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, "err-invalid-id-format")
		return
	}

	author, err := a.dbQueries.GetAuthor(r.Context(), int32(id))
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			writeJSONResponse(w, http.StatusInternalServerError, errors.New("err-author-id-not-found"))
			return
		}
		writeJSONResponse(w, http.StatusInternalServerError, errors.New("err-failed-to-get-author"))
		return
	}

	//For now change it to dummy later take from payload
	authorParams := database.UpdateAuthorParams{
		ID:   author.ID,
		Name: "Dummy",
		Bio:  author.Bio,
	}
	_, err = a.dbQueries.UpdateAuthor(r.Context(), authorParams)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			writeJSONResponse(w, http.StatusInternalServerError, errors.New("err-author-id-not-found"))
			return
		}
		writeJSONResponse(w, http.StatusInternalServerError, errors.New("err-failed-to-get-author"))
		return
	}
	writeJSONResponse(w, http.StatusNoContent, "")
}
