package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) RegisterRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/ping", a.Ping).Methods("GET")

	//Designing routes is one of my favorite thing

	//1. plural Nouns
	//2. hierarchical structures
	//3. consistency: includes naming, request, response structure...
	//4. versioning: optional but require for some usecase
	//5. query param: filtering, pagination, sorting
	//6. and most imp avoid verbs :) profits :)

	//will follow strict response status like 201 created, 204 no content
	//thanks OCI for teaching me this... :)

	//Author routes
	r.HandleFunc("/authors", a.listAuthorsHandler).Methods("GET")
	r.HandleFunc("/authors", a.createAuthorHandler).Methods("POST")

	r.HandleFunc("/authors/{id}", a.getAuthorHandler).Methods("GET")
	r.HandleFunc("/authors/{id}", a.Ping).Methods("PUT")
	r.HandleFunc("/authors/{id}", a.Ping).Methods("DELETE")

	//Book routes
	r.HandleFunc("/books", a.Ping).Methods("GET")
	r.HandleFunc("/books", a.Ping).Methods("POST")

	r.HandleFunc("/books/{id}", a.Ping).Methods("GET")
	r.HandleFunc("/books/{id}", a.Ping).Methods("PUT")
	r.HandleFunc("/books/{id}", a.Ping).Methods("DELETE")

	//Author-Book Relationship
	r.HandleFunc("/authors/{id}/books", a.Ping).Methods("GET")
	r.HandleFunc("/authors/{id}/books", a.Ping).Methods("POST")

	return r
}
