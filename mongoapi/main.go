package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Thenameisafsal/mongoapi/router"
)

func main() {
	fmt.Println("Starting the server")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listening at port 4000")
}
