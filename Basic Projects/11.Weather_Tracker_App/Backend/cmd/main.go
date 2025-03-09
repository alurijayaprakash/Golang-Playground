package main

import (
	"WeatherTrackerApp/internal/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// var r = mux.NewRouter() // router by mux
	router.InitRouter()
	fmt.Println("Starting server ...!")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
