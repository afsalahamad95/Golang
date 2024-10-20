package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/api/students", AllStudents).Methods("GET")
	return router

}

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var students []Student

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to my page</h1>"))
}

func AllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for _, student := range students {
		json.NewEncoder(w).Encode(student)
	}
}

func main() {
	students = append(students, Student{Id: 97, Name: "afsal"})
	students = append(students, Student{Id: 100, Name: "mugesh"})
	log.Fatal(http.ListenAndServe(":4000", Router()))
}
