package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// seeding of data
	courses = append(courses, Course{CourseId: "2", Name: "Go", Price: 999, Rating: 4.8, Author: &Author{FullName: "John wick", Website: "https://golang.dev"}})
	courses = append(courses, Course{CourseId: "3", Name: "Flutter", Price: 999, Rating: 4.8, Author: &Author{FullName: "John wick", Website: "https://golang.dev"}})

	// routing
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET") // pass the id in params through /{} syntax
	router.HandleFunc("/delete-all", deleteAllCourses).Methods("DELETE")
	router.HandleFunc("/create", createCourse).Methods("POST")
	router.HandleFunc("/update/{id}", updateOneCourse).Methods("PUT")

	// listen to port
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
	// return c.CourseId == "" && c.Name == ""
	return c.Name == ""
}

// controllers -> separate fils

// define a route
func serveHome(w http.ResponseWriter, r *http.Request) {
	// send response
	w.Write([]byte("<h1>hello world</h1>")) // byte slice
}

// implement CRUD operations

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

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty body detected - please send some data")
		return
	}
	// if data sent like {}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data detected")
		return
	}

	// generate unique id, convert to string
	// append course to courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100)) // convert to string
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// loop and get id, remove the id and add new id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// remove id
			courses = append(courses[:index], courses[index+1:]...) // variadic operation, so use ...
			var course Course
			json.NewDecoder(r.Body).Decode(&course) // write current data to variable
			course.CourseId = params["id"]          // update the id as per user data
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course) // notify the operation
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // delete the course
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	courses = []Course{} // reassign
	json.NewEncoder(w).Encode("delete successful")
}
