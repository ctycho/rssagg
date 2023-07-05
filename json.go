package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if (err != nil) {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if (code > 499) {
		log.Fatal("Responding with 5XX status code")
	}
	type ErrorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, ErrorResponse{
		Error: msg,
	})
}