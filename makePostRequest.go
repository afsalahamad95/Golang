package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	PostRequestJson()
}

func PostRequestJson() {
	const myUrl = "http://localhost:8000/post"
	// create a fake json payload
	requestBody := strings.NewReader(`
	{
		"course_name":"golang",
		"price":999
	}
	`)
	// application/json is the content type
	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println(string(content))
}
