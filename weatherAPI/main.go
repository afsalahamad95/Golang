package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const apikey = "http://api.weatherapi.com/v1/current.json?key=6040383dfe83421a88684557242010&q=India&aqi=no"

func parser(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(apikey)
	if err != nil {
		fmt.Println("error fetching data", err)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var res Responsew
	err = json.Unmarshal(body, &res)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	defer resp.Body.Close()
	return
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", parser).Methods("GET")
	http.ListenAndServe(":4000", router)
}

type Responsew struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
	} `json:"current"`
}
