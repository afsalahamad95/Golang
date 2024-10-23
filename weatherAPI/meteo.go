package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	Latitude             float64       `json:"latitude"`
	Longitude            float64       `json:"longitude"`
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
	Latitude             float64       `json:"latitude" bson:"latitude"`
	Longitude            float64       `json:"longitude" bson:"longitude"`
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

func refreshDB() {
	resp, err := http.Get(api)
	if err != nil {
		log.Fatal(err)
	}
	var Apiresponse apiresponse
	parse, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(parse, &Apiresponse)
	fmt.Println("update starting")
	FinalResponse := convertToFinalResponse(Apiresponse)
	insertWeather(FinalResponse)
	fmt.Println("Writing to database")
	fmt.Println("write success")
	defer resp.Body.Close()
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
	}
}

var collection *mongo.Collection

func insertWeather(data finalresponse) {
	// deleteAllRecords()
	status, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		fmt.Println("error occured during insertion")
		return
	}
	fmt.Println("updated successfully with id:", status.InsertedID)
}

func deleteAllRecords() {
	res, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted", res.DeletedCount, "files")
}
func autoupdatedb() {
	ticker := time.NewTicker(20 * time.Minute)
	defer ticker.Stop()
	for {
		<-ticker.C
		// non blocking ticker
		refreshDB()
		deleteOldRecords()
		fmt.Println("auto update triggered")
	}
}

type LocationDetails struct {
	Lat string
	Lon string
}

func getFromDb(w http.ResponseWriter, r *http.Request) {
	var location LocationDetails
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&location)
	latitude, _ := strconv.ParseFloat(location.Lat, 64)
	longitude, _ := strconv.ParseFloat(location.Lon, 64)
	findOptions := options.FindOne()
	updateApi(location.Lat, location.Lon)
	var res finalresponse
	err := collection.FindOne(context.Background(), bson.M{"latitude": latitude, "longitude": longitude}, findOptions).Decode(&res)
	if err == mongo.ErrNoDocuments {
		refreshDB()
	}
	err = collection.FindOne(context.Background(), bson.M{"latitude": latitude, "longitude": longitude}, findOptions).Decode(&res)
	json.NewEncoder(w).Encode(res)
}

func main() {
	// let's connect to mongodb
	fmt.Println("Connecting to MongoDB")
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Connection success!")
	go autoupdatedb() // auto update db as per time
	router := Router()
	http.ListenAndServe(":4000", router)
}

func updateApi(a, b string) {
	// update api with new latitude and longitude
	api = fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current=temperature_2m,relative_humidity_2m,weather_code,surface_pressure,wind_speed_10m&hourly=temperature_2m,relative_humidity_2m,weather_code,surface_pressure,visibility,wind_speed_10m", a, b)
}

func deleteOldRecords() {
	twentyMinutesAgo := time.Now().Add(-20 * time.Minute)
	filter := bson.M{"recordtime": bson.M{"$lt": twentyMinutesAgo}}
	status, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(status.DeletedCount, "files deleted")
}
