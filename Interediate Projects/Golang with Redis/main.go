package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server")

	http.HandleFunc("/api", Handler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler Triggered")
}
