package app

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data any) {

	if statusCode == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{}

	//Following structure {"data":"DATA"} if success and if error simple {"error": "err-actual-error-unique-error"} \
	//all error must start from "err-" it will distinguish error from defined or internal package

	if err, ok := data.(error); ok {
		response.Error = err.Error()
	} else {
		response.Data = data
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("failed to write JSON response %v", err)
	}
}
