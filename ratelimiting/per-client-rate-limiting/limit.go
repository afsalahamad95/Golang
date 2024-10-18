package main

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func RateLimiter(next http.HandlerFunc) http.Handler {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}
	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	// Periodically remove inactive clients
	go func() {
		for {
			time.Sleep(time.Minute) // Run cleanup every minute
			mu.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip) // Remove clients not seen for over 3 minutes
				}
			}
			mu.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get client IP address
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		mu.Lock()
		if _, found := clients[ip]; !found {
			clients[ip] = &client{limiter: rate.NewLimiter(2, 4)} // 2 requests per 4 seconds
		}
		clients[ip].lastSeen = time.Now() // Update last seen time

		// Check if the request is allowed
		if !clients[ip].limiter.Allow() {
			mu.Unlock()
			w.WriteHeader(http.StatusTooManyRequests) // Return 429 status
			msg := message{Status: "Too many requests", Body: "Server busy"}
			json.NewEncoder(w).Encode(&msg)
			return
		}
		mu.Unlock()

		// Call the next handler if rate limit is respected
		next(w, r)
	})
}
