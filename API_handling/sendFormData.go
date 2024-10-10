package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	performPostFormRequest()
}

func performPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"
	// create form data
	data := url.Values{}
	data.Add("firstname", "afsal")
	data.Add("lastname", "ahamad")
	data.Add("age", "25")
	response, err := http.PostForm(myUrl, data) // url encoded
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(string(content))
}
