package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const connectionString = "mongodb+srv://golang:I@mAfsal05@cluster0.1bz0k.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

// important
var collection *mongo.Collection // pointer

// connect with mongodb
// runs only once during initialisation
func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection success!")
	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("collection instance ready")
}
