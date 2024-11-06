package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventNotification struct {
	Content string    `json:"content" bson:"content"`
	Timings time.Time `json:"timings" bson:"timings"`
}

type LeaveNotification struct {
	Content string    `json:"content" bson:"content"`
	Timings time.Time `json:"timings" bson:"timings"`
	Author  string    `json:"author" bson:"author"`
}

type GeneralNotification struct {
	Content string    `json:"content" bson:"content"`
	Timings time.Time `json:"timings" bson:"timings"`
	Author  string    `json:"author" bson:"author"`
}

func viperConfigure(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("error reading the env file")
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Println("Invalid configuration")
	}
	return value
}

var collection *mongo.Collection

func connect_mongo() {
	connectionString := viperConfigure("CONNECTION_STRING")
	dbname := viperConfigure("DBNAME")
	colname := viperConfigure("COLLECTIONNAME")
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println("an error occurred during db configuration")
	}
	collection = client.Database(dbname).Collection(colname)
	log.Println("collection obtained")
}

func addEventNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	defer r.Body.Close()
	var res EventNotification
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		log.Println("error occurred while decoding the event notification")
	}
	content := res.Content
	res.Timings = time.Now()
	timings := res.Timings
	notification := EventNotification{Content: content, Timings: timings}
	result, err := collection.InsertOne(context.Background(), notification)
	if err != nil {
		log.Println("an error occurred during the insertion of event data")
		return
	}
	log.Println("event notification data insert successful with id", result.InsertedID)
	json.NewEncoder(w).Encode(res)
}

func addLeaveNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var res LeaveNotification
	error := json.NewDecoder(r.Body).Decode(&res)
	defer r.Body.Close()
	if error != nil {
		log.Println("an error occurred when decoding the leave notification")
	}
	res.Timings = time.Now()
	timings := res.Timings
	content := res.Content
	author := res.Author
	document := LeaveNotification{Content: content, Timings: timings, Author: author}
	result, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		log.Println("An error occurred during the insertion of leave data")
		return
	}
	log.Println("leave notification insert success with id", result.InsertedID)
	json.NewEncoder(w).Encode(res)
}

func addGeneralNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var result GeneralNotification
	json.NewDecoder(r.Body).Decode(&result)
	defer r.Body.Close()
	result.Timings = time.Now()
	timings := result.Timings
	content := result.Content
	author := result.Author
	document := GeneralNotification{Content: content, Timings: timings, Author: author}
	res, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		log.Println("An error occurred during the insertion of general notification")
		return
	}
	log.Println("general notification insertion success with id", res.InsertedID)
	json.NewEncoder(w).Encode(result)
}

func main() {
	connect_mongo() // connect to the db
	router := mux.NewRouter()
	router.HandleFunc("/insertGeneral", addGeneralNotification)
	router.HandleFunc("/insertEvent", addEventNotification)
	router.HandleFunc("/insertLeave", addLeaveNotification)
	log.Fatal(http.ListenAndServe(":4000", router))
}
