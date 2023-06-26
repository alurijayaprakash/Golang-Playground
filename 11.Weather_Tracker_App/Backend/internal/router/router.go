package router

import (
	"WeatherTrackerApp/internal/app"
	"encoding/json"
	"fmt"
	"net/http"
)

const APIKEY = "b3740c800de5629c9096b5f2068d74b5"

func InitRouter() {
	// var r = mux.NewRouter() // router by mux
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/track", TrackHandler) // localhost:8080/track?q=Mumbai&appid=1f485364c2c92f74eaca9d0bb5768b5f
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func TrackHandler(w http.ResponseWriter, r *http.Request) {
	CityName := r.URL.Query().Get("q") // city name
	// Appid := r.URL.Query().Get("appid") // appid // // localhost:8080/track?q=Mumbai&appid=1f485364c2c92f74eaca9d0bb5768b5f
	Appid := APIKEY // // localhost:8080/track?q=Mumbai
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(CityName, Appid)
	data, err := app.QueryOpenWeatherAPI(CityName, Appid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
