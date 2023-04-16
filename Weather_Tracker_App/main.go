package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// const APIKEY = "b3740c800de5629c9096b5f2068d74b5"
const APIKEY_URL = "http://api.openweathermap.org/data/2.5/weather"

type WeatherData struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// http://api.openweathermap.org/data/2.5/weather?q=Hyderabad&appid=1f485364c2c92f74eaca9d0bb5768b5f
// http://api.openweathermap.org/data/2.5/weather&appid=b3740c800de5629c9096b5f2068d74b5?q=Hyderabad
func QueryOpenWeatherAPI(city string, Appid string) (WeatherData, error) {
	QueryString := fmt.Sprintf(APIKEY_URL + "?q=" + city + "&appid=" + Appid)
	fmt.Println(QueryString)
	resp, err := http.Get(QueryString)
	if err != nil {
		return WeatherData{}, err
	}

	defer resp.Body.Close()

	var Wd WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&Wd); err != nil {
		return WeatherData{}, err
	}
	fmt.Println(Wd)
	return Wd, nil
}

func trackHandler(w http.ResponseWriter, r *http.Request) {
	CityName := r.URL.Query().Get("q")  // city name
	Appid := r.URL.Query().Get("appid") // appid

	fmt.Println(CityName, Appid)
	data, err := QueryOpenWeatherAPI(CityName, Appid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	var r = mux.NewRouter() // router by mux

	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/track", trackHandler).Methods("GET") // localhost:8080/track?q=Mumbai&appid=1f485364c2c92f74eaca9d0bb5768b5f
	fmt.Println("Starting server ...!")

	log.Fatal(http.ListenAndServe(":8080", r))
}
