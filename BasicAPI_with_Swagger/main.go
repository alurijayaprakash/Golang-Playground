package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "swaggerapi/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Item represents a simple data model
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetAllItems returns a list of items
// @Summary Get a list of items
// @Description Get all items
// @Produce json
// @Success 200 {array} Item
// @Router /items [get]
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	items := []Item{
		{ID: "1", Name: "Item 1"},
		{ID: "2", Name: "Item 2"},
		// Add more items as needed
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// GetItemByID returns details of a specific item by ID
// @Summary Get item by ID
// @Description Get details of a specific item
// @Produce json
// @Param id path string true "Item ID"
// @Success 200 {object} Item
// @Router /items/{id} [get]
func GetItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]

	item := Item{ID: itemID, Name: fmt.Sprintf("Item %s", itemID)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

// @title BASICAPI_with_Swagger
// @version 1.0
// @description This is a sample REST API with Swagger documentation.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/api/v1/items", GetAllItems).Methods("GET")
	router.HandleFunc("/api/v1/items/{id}", GetItemByID).Methods("GET")

	// Serve Swagger documentation
	router.PathPrefix("/swagger").Handler(httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))

	// Start the server
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
