package main

import (
	"fmt"
	"log"
	"net/http"
	"stocksapi/handlers"
	"stocksapi/models"
	"stocksapi/service"
	"stocksapi/store"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, World!")

	repo := store.NewStockRepository()
	service := service.NewStockService(repo)
	handler := handlers.NewStockHandler(service)
	router := setupRouter(handler)
	ser := http.Server{
		Addr:    ":8888",
		Handler: router,
	}
	log.Println("Server will start")
	err := ser.ListenAndServe()
	if err != nil {
		panic("HTTP Server Error")
	}
}

func setupRouter(handler models.StockHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/stocks", handler.ListStocksHandler).Methods(http.MethodGet)
	router.HandleFunc("/stocks/{id}", handler.GetStockHandler).Methods(http.MethodGet)
	router.HandleFunc("/stocks", handler.CreateStockHandler).Methods(http.MethodPost)
	router.HandleFunc("/stocks/{id}", handler.UpdateStockHandler).Methods(http.MethodPut)
	router.HandleFunc("/stocks/{id}", handler.DeleteStockHandler).Methods(http.MethodDelete)
	return router
}
