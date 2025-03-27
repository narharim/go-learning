package app

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := Response{Data: data}

	if errMsg != "" {
		resp.Error = errMsg
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to write JSON response %v", err)
	}
}

func HandleAppError(w http.ResponseWriter, err error) {
	WriteJSONResponse(w, http.StatusInternalServerError, &Response{
		Error: err.Error(),}
	)
}
