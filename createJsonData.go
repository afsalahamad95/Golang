package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"coursename"` // the third parameter is a custom name parameter that will be displayed in the json fils
	Price    int
	password string   `json:"-"`              // if you give - as the custom name, the value will not be displayed
	tags     []string `json:"tags,omitempty"` // empty values are not displayed because of omitempty
}

func main() {
	decodeJson()
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
	// the second arg is a prefix arg, best to leave it empty, otherwise in every line there will be the prefix string which you enter there
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

// consume a json data
func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
			"Name": "React",
			"Price": 999
	}
	`)

	var someCourse Course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json valid!")
		json.Unmarshal(jsonDataFromWeb, &someCourse) // result is stored to someCourse var here
		fmt.Printf("%#v\n", someCourse)
	} else {
		fmt.Println("Invalid json!")
	}
}
