package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var api string = "https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current=temperature_2m,relative_humidity_2m,weather_code,surface_pressure,wind_speed_10m&hourly=temperature_2m,relative_humidity_2m,weather_code,surface_pressure,visibility,wind_speed_10m"

const connectionString = "mongodb+srv://afsal:afsal12345@weatherapi.7xuwg.mongodb.net/?retryWrites=true&w=majority&appName=weatherAPI"
const dbName = "weather"
const colName = "weatherdata"

var mut sync.Mutex

type CurrentUnits struct {
	Time               string `json:"time"`
	Interval           string `json:"interval"`
	Temperature2m      string `json:"temperature_2m"`
	RelativeHumidity2m string `json:"relative_humidity_2m"`
	WeatherCode        string `json:"weather_code"`
	SurfacePressure    string `json:"surface_pressure"`
	WindSpeed10m       string `json:"wind_speed_10m"`
}

type Current struct {
	Time               string  `json:"time"`
	Interval           int     `json:"interval"`
	Temperature2m      float64 `json:"temperature_2m"`
	RelativeHumidity2m int     `json:"relative_humidity_2m"`
	WeatherCode        int     `json:"weather_code"`
	SurfacePressure    float64 `json:"surface_pressure"`
	WindSpeed10m       float64 `json:"wind_speed_10m"`
}

type THourly_units struct {
	Time               string `json:"time"`
	Temperature2m      string `json:"temperature_2m"`
	RelativeHumidity2m string `json:"relative_humidity_2m"`
	WeatherCode        string `json:"weather_code"`
	SurfacePressure    string `json:"surface_pressure"`
	Visibility         string `json:"visibility"`
	WindSpeed10m       string `json:"wind_speed_10m"`
}

type apiresponse struct {
	Latitude             string        `json:"latitude"`
	Longitude            string        `json:"longitude"`
	GenerationtimeMs     float64       `json:"generationtime_ms"`
	UtcOffsetSeconds     int           `json:"utc_offset_seconds"`
	Timezone             string        `json:"timezone"`
	TimezoneAbbreviation string        `json:"timezone_abbreviation"`
	Elevation            float64       `json:"elevation"`
	CurrentUnits         CurrentUnits  `json:"current_units"`
	Current              Current       `json:"current"`
	Hourly_units         THourly_units `json:"hourly_units"`
	Hourly               HourlyData    `json:"hourly"`
}
type finalresponse struct {
	Latitude             string        `json:"latitude" bson:"latitude"`
	Longitude            string        `json:"longitude" bson:"longitude"`
	GenerationtimeMs     float64       `json:"generationtime_ms" bson:"generationtime_ms"`
	UtcOffsetSeconds     int           `json:"utc_offset_seconds" bson:"utc_offset_seconds"`
	Timezone             string        `json:"timezone" bson:"timezone"`
	TimezoneAbbreviation string        `json:"timezone_abbreviation" bson:"timezone_abbreviation"`
	Elevation            float64       `json:"elevation" bson:"elevation"`
	CurrentUnits         CurrentUnits  `json:"current_units" bson:"current_units"`
	Current              Current       `json:"current" bson:"current"`
	Hourly_units         THourly_units `json:"hourly_units" bson:"hourly_units"`
	Hourly               HourlyData    `json:"hourly" bson:"hourly"`
	RecordTime           time.Time     `json:"recordtime" bson:"recordtime"`
	City                 string        `json:"city" bson:"city"`
}
type HourlyData struct {
	Time               []string  `json:"time"`
	Temperature2m      []float64 `json:"temperature_2m"`
	RelativeHumidity2m []int     `json:"relative_humidity_2m"`
	WeatherCode        []int     `json:"weather_code"`
	SurfacePressure    []float64 `json:"surface_pressure"`
	Visibility         []float64 `json:"visibility"`
	WindSpeed10m       []float64 `json:"wind_speed_10m"`
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", getFromDb)
	return router
}

func refreshDB(city string, latitude string, longitude string) finalresponse {
	if city == "" || latitude == "" || longitude == "" {
		log.Println("empty fields detected in refresh")
	}
	resp, err := http.Get(api)
	if err != nil {
		fmt.Println("an error occurred when fetching the data", err)
	}
	var Apiresponse apiresponse
	parse, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(parse, &Apiresponse)
	fmt.Println("update starting")
	FinalResponse := convertToFinalResponse(Apiresponse)
	FinalResponse.City = city
	FinalResponse.Latitude = latitude
	FinalResponse.Longitude = longitude
	fmt.Println("checking", FinalResponse.Latitude, FinalResponse.Longitude)
	insertRes := insertWeather(FinalResponse)
	// fmt.Println(insertRes, "status from insertion")
	fmt.Println("Writing to database")
	fmt.Println("write success")
	defer resp.Body.Close()
	return insertRes
}
func convertToFinalResponse(Apiresponse apiresponse) finalresponse {
	return finalresponse{
		Latitude:             Apiresponse.Latitude,
		Longitude:            Apiresponse.Longitude,
		GenerationtimeMs:     Apiresponse.GenerationtimeMs,
		UtcOffsetSeconds:     Apiresponse.UtcOffsetSeconds,
		Timezone:             Apiresponse.Timezone,
		TimezoneAbbreviation: Apiresponse.TimezoneAbbreviation,
		Elevation:            Apiresponse.Elevation,
		CurrentUnits:         Apiresponse.CurrentUnits,
		Current:              Apiresponse.Current,
		Hourly_units:         Apiresponse.Hourly_units,
		Hourly:               Apiresponse.Hourly,
		RecordTime:           time.Now(),
		City:                 "Coimbatore",
	}
}

var collection *mongo.Collection

func insertWeather(data finalresponse) finalresponse {
	// deleteAllRecords()
	status, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			fmt.Println("insert operation aborted due to duplicate entry, returning existing entry")
			var existingRec finalresponse
			search_city := data.City
			err := collection.FindOne(context.Background(), bson.M{"city": search_city}).Decode(existingRec)
			if err == nil {
				return existingRec
			}
		}
		return data
	}
	fmt.Println("updated successfully with id:", status.InsertedID)
	return data
}

func deleteAllRecords() {
	res, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		fmt.Println("error occurred when deleting records")
	}
	fmt.Println("Deleted", res.DeletedCount, "files")
}
func autodeletedb() {
	ticker := time.NewTicker(20 * time.Minute)
	defer ticker.Stop()
	for {
		<-ticker.C
		// non blocking ticker
		// refreshDB("Coimbatore", "11.0787684", "77.0370419")
		deleteOldRecords()
		fmt.Println("auto delete triggered")
	}
}

type LocationDetails struct {
	Lat  string
	Lon  string
	City string
}

func getFromDb(w http.ResponseWriter, r *http.Request) {
	var location LocationDetails
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		log.Println("Error decoding the coordinates - warning")
	}
	defer r.Body.Close()
	latitude := location.Lat
	longitude := location.Lon
	city := location.City
	if city == "" || latitude == "" || longitude == "" {
		log.Println("empty fields detected")
		json.NewEncoder(w).Encode("empty fields detected")
	} else {
		fmt.Println("the city in query is:", city)
		fmt.Println(latitude, longitude)
		updateApi(latitude, longitude)
		findOptions := options.FindOne()
		var res finalresponse
		err = collection.FindOne(context.Background(), bson.M{"city": city}, findOptions).Decode(&res)
		if err == mongo.ErrNoDocuments {
			fmt.Println("call forwarding to api, since not found in database")
			res = refreshDB(city, latitude, longitude)
		} else {
			fmt.Println("found match in db")
		}
		json.NewEncoder(w).Encode(res)
	}
}

func main() {
	// let's connect to mongodb
	fmt.Println("Connecting to MongoDB")
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("error occurred during connection with database")
	}
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Connection success!")
	fmt.Println("indexing")
	// create index
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			bson.E{Key: "latitude", Value: 1},  // ascending order
			bson.E{Key: "longitude", Value: 1}, // ascending order
			bson.E{Key: "city", Value: 1},
		}, Options: options.Index().SetUnique(true),
	}
	_, err = collection.Indexes().CreateOne(context.Background(), indexModel)
	// deleteAllRecords()
	if err != nil {
		fmt.Println("indexing error occurred", err)
	} else {
		fmt.Println("indexing complete")
	}
	go autodeletedb() // auto update db as per time
	router := Router()
	http.ListenAndServe(":4000", router)
}

func updateApi(a, b string) {
	// update api with new latitude and longitude
	mut.Lock()
	api = fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current=temperature_2m,relative_humidity_2m,weather_code,surface_pressure,wind_speed_10m&hourly=temperature_2m,relative_humidity_2m,weather_code,surface_pressure,visibility,wind_speed_10m", a, b)
	mut.Unlock()
}

func deleteOldRecords() {
	fmt.Println("deleting old records now")
	twentyMinutesAgo := time.Now().Add(-20 * time.Minute)
	filter := bson.M{"recordtime": bson.M{"$lt": twentyMinutesAgo}}
	status, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(status.DeletedCount, "files deleted")
}
