package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getOneCourse)
	log.Fatal(http.ListenAndServe(":4000", router))
}

// course file model -> separate file
type Course struct {
	CourseId string  `json:"id"`
	Name     string  `json:"name"`
	Price    int     `json:"price"`
	Rating   float64 `json:"rating"`
	Author   *Author `json:"author"` // *author is a type pointer
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware, helpers -> separate file
func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.Name == ""
}

// define a route
func serveHome(w http.ResponseWriter, r *http.Request) {
	// send response
	w.Write([]byte("<h1>hello world</h1>")) // byte slice
}

// return courses in json format
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // data is sent in json format
	json.NewEncoder(w).Encode(courses)                 // get the db in json format and writes it to w
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	params := mux.Vars(r) // get variables
	// fmt.Println(params)
	// fmt.Printf("%T", params)

	// loop through courses, find the matching id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}
