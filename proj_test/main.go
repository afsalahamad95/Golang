package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	api := "http://192.168.1.87:3000/get"
	// res, err := http.Post(api, "application/json", strings.NewReader("Coimbatore"))
	res, err := http.Get(api)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	defer res.Body.Close()
}
