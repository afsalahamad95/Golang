package main

import (
	"encoding/json"
	"log"
	"net/http"

	tollbooth "github.com/didip/tollbooth/v7"
)

type message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endPointHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	msg := message{Status: "OK", Body: "Hello , auth success"}
	err := json.NewEncoder(w).Encode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	msg := message{Status: "failed", Body: "sorry , auth failed"}
	jsonmsg, _ := json.Marshal(msg)
	tlbthLimiter := tollbooth.NewLimiter(1, nil)
	tlbthLimiter.SetMessageContentType("application/json")
	tlbthLimiter.SetMessage(string(jsonmsg))
	http.Handle("/ping", tollbooth.LimitFuncHandler(tlbthLimiter, endPointHandler))
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
