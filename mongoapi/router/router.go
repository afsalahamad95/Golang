package router

import (
	"github.com/Thenameisafsal/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMyMovies).Methods("GET")
	router.HandleFunc("/api/create", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/delete/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/delete-all", controller.DeleteAllMovies).Methods("DELETE")
	return router
}

func main() {

}
