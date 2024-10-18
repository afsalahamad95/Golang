package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

// handle what happens when a request is made
func endPointHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	msg := message{Status: "success", Body: "hello"}
	err := json.NewEncoder(w).Encode(&msg)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	http.Handle("/ping", RateLimiter(endPointHandler))
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Println("There was an error", err)
	}
}
