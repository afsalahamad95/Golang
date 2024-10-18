package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimiter(next func(http.ResponseWriter, *http.Request)) http.Handler {
	limiter := rate.NewLimiter(2, 4) // 2 requests every 4 seconds
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			msg := message{Status: "too many requests", Body: "please try again later"}
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&msg)
			return
		} else {
			next(w, r)
		}
	})
}
