package app

import (
	"WeatherTrackerApp/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// const APIKEY = "b3740c800de5629c9096b5f2068d74b5"
const APIKEY_URL = "http://api.openweathermap.org/data/2.5/weather"

// http://api.openweathermap.org/data/2.5/weather?q=Hyderabad&appid=1f485364c2c92f74eaca9d0bb5768b5f
// http://api.openweathermap.org/data/2.5/weather&appid=b3740c800de5629c9096b5f2068d74b5?q=Hyderabad
func QueryOpenWeatherAPI(city string, Appid string) (models.WeatherData, error) {
	QueryString := fmt.Sprintf(APIKEY_URL + "?q=" + city + "&appid=" + Appid)
	fmt.Println(QueryString)
	resp, err := http.Get(QueryString)
	if err != nil {
		return models.WeatherData{}, err
	}

	defer resp.Body.Close()

	var Wd models.WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&Wd); err != nil {
		return models.WeatherData{}, err
	}
	// fmt.Println(Wd)
	return Wd, nil
}
