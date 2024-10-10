package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string
	Price    int
	password string
	tags     []string
}

func main() {
	encodeJson()
}

func encodeJson() {
	// slice of structs
	courses := []Course{
		{"React", 999, "12345", []string{"Dev", "dev"}},
		{"Flutter", 999, "12345", []string{"Dev", "dev"}},
		{"Go", 999, "12345", []string{"Dev", "dev"}},
		{"JS", 999, "12345", []string{"Dev", "dev"}},
	}

	// package the data as json data
	// pass an interface here
	finalJson, err := json.MarshalIndent(courses, "", "\t") // we're using marshal indent for making the result json readable
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
