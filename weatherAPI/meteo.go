package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const a float64 = 52.52
//
//	const b float64 = 13.41
//	var api string = fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.2f&longitude=%.2f&hourly=temperature_2m", a, b)
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
	Hourly       HourlyData    `json:"hourly"`
}
type HourlyData struct {
	Time          []string  `json:"time"`
	Temperature2m []float64 `json:"temperature_2m"`
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", getFromDb).Methods("GET")
	return router
}

func controller() {
	resp, err := http.Get(api)
	if err != nil {
		log.Fatal(err)
	}
	var Apiresponse apiresponse
	parse, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(parse, &Apiresponse)
	fmt.Println(parse)
	fmt.Println("update starting")
	insertWeather(Apiresponse)
	fmt.Println("Writing to database")
	fmt.Println("write success")
	defer resp.Body.Close()
}

var collection *mongo.Collection
var lastInsertedId primitive.ObjectID = primitive.NilObjectID

func insertWeather(data apiresponse) {
	deleteAllRecords()
	status, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		fmt.Println("error occured")
		return
	}
	fmt.Println("updated successfully with id:", status.InsertedID)
	lastInsertedId = status.InsertedID.(primitive.ObjectID)
}

func deleteAllRecords() {
	res, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted", res.DeletedCount, "files")
}
func autoupdatedb() {
	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()
	for {
		<-ticker.C
		controller()
		fmt.Println("auto update triggered")
	}
}
func getFromDb(w http.ResponseWriter, r *http.Request) {
	if lastInsertedId == primitive.NilObjectID {
		fmt.Println("No data to fetch,fetching from api")
		controller()
	} else {
		filter := bson.M{"_id": lastInsertedId}
		var result apiresponse
		err := collection.FindOne(context.Background(), filter).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(result)
	}
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
	go autoupdatedb() // auto update db as per time
	time.Sleep(1 * time.Second)
	router := Router()
	http.ListenAndServe(":4000", router)
}
