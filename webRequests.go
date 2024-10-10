package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The type is %T", response) // we get a *http.Response which is a pointer to the response

	defer response.Body.Close() // we've to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(databytes)
	fmt.Println(string(databytes))
}
