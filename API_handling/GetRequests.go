package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	performGetRequest()
}

func performGetRequest() {
	const myUrl = "http://localhost:8000/get"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code", response.StatusCode)
	fmt.Println("Length of content:", response.ContentLength)
	// to read the response as string without using string() to convert the bytes to string, use:
	// method 1:
	var responseString strings.Builder
	// this string is same for both methods
	res, err := ioutil.ReadAll(response.Body)
	// method 1
	byteCount, _ := responseString.Write(res) // get the response
	fmt.Println("byte count is:", byteCount)  // gives length
	fmt.Println(responseString.String())      // prints content

	// method 2:
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(res))
}
