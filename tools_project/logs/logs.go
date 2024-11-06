package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Logs struct {
	Checkin  time.Time `json:"checkin" bson:"checkin"`
	Checkout time.Time `json:"checkout" bson:"checkout"`
	Username string    `json:"username" bson:"username"`
}

var collection *mongo.Collection

func insertLog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var res Logs
	error := json.NewDecoder(r.Body).Decode(&res)
	if error != nil {
		log.Println("an error occurred during the log data insert operation")
	}
	result, err := collection.InsertOne(context.Background(), res)
	if err != nil {
		log.Println("error occurred during the insertion of log data")
	}
	log.Println("insertion was successful with log id", result.InsertedID)
	json.NewEncoder(w).Encode(res)
}

const connectionString = "mongodb://localhost:27017/"
const dbname = "logs"
const colname = "logdata"

func connect_mongo() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println("error occurred when creating the collection")
	}
	collection = client.Database(dbname).Collection(colname)
	log.Println("database is connected")
}

func main() {
	connect_mongo()
	indexModel := mongo.IndexModel{Keys: bson.D{
		bson.E{Key: "Username", Value: 1},
		bson.E{Key: "Checkin", Value: 1},
		bson.E{Key: "Checkout", Value: 1},
	}}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Println("indexing error occurred")
	}
	log.Println("Indexing complete")
	router := mux.NewRouter()
	router.HandleFunc("/insertlogs", insertLog)
	log.Fatal(http.ListenAndServe(":4000", router))
}
