package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const api = "https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&hourly=temperature_2m"
const connectionString = "mongodb+srv://afsal:afsal12345@weatherapi.7xuwg.mongodb.net/?retryWrites=true&w=majority&appName=weatherAPI"
const dbName = "weather"
const colName = "weatherdata"

type THourly_units struct {
	Time           string `json:"time"`
	Temperature_2m string `json:"temperature_2m"`
}
type apiresponse struct {
	Latitude     float64       `json:"latitude"`
	Longitude    float64       `json:"longitude"`
	Hourly_units THourly_units `json:"hourly_units"`
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller).Methods("GET")
	return router
}
func controller(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(api)
	if err != nil {
		log.Fatal(err)
	}
	var Apiresponse apiresponse
	parse, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(parse, &Apiresponse)
	fmt.Println(parse)
	fmt.Println("Insert starting")
	insertWeather(Apiresponse)
	json.NewEncoder(w).Encode(Apiresponse)
	defer resp.Body.Close()
}

var collection *mongo.Collection

func insertWeather(data apiresponse) {
	status, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		fmt.Println("error occured")
		return
	}
	fmt.Println("inserted successfully with id:", status.InsertedID)
}
func main() {
	// let's connect to mongodb
	fmt.Println("Connecting to MongoDB")
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Connection success!")
	router := Router()
	http.ListenAndServe(":4000", router)
}
