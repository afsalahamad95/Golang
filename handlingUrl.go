package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://google.com"

func main() {
	// parse the url
	response, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
	fmt.Println(response.Scheme)   // which scheme - http, https etc.
	fmt.Println(response.Host)     // host name
	fmt.Println(response.Port())   // get the port number
	fmt.Println(response.RawQuery) // the query sent, anything after the ?
	// get the query parameters
	qparams := response.Query()
	fmt.Println(qparams)
	fmt.Printf("Type is %T", qparams) // stored in key value pairs

	// in case u don't know url, get it like this:
	// make sure to pass the reference
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "google.com",
		Path:    "/home",
		RawPath: "user=home",
	}
	resUrl := partsOfUrl.String()
	fmt.Println(resUrl)
}
