package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hey bub")
	greet()
	r := mux.NewRouter()                        // create the router
	r.HandleFunc("/", serveHome).Methods("GET") // handle with the function defined below
	// the methods("get") will specify that this method will accept only get requests, or whatever is specified
	log.Fatal(http.ListenAndServe(":4000", r)) // we're listening on port 4000 and serving request r
	// use log.fatal to log any errors
}

func greet() {
	fmt.Println("Yo, what's up nigga?")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	// r will have the parameters and all
	// w is used to send the response
	w.Write([]byte("<h1>Yo Sup nigga?</h1>"))
}

// after go build, run the file, and localhost:4000 will have the content we've written here
