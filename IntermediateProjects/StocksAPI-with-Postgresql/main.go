package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"stocksapi/handlers"
	"stocksapi/models"
	"stocksapi/service"
	"stocksapi/store"
	"syscall"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

func setupRouter(handler models.StockHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/stocks", handler.ListStocksHandler).Methods(http.MethodGet)
	router.HandleFunc("/stocks/{id}", handler.GetStockHandler).Methods(http.MethodGet)
	router.HandleFunc("/stocks", handler.CreateStockHandler).Methods(http.MethodPost)
	router.HandleFunc("/stocks/{id}", handler.UpdateStockHandler).Methods(http.MethodPut)
	router.HandleFunc("/stocks/{id}", handler.DeleteStockHandler).Methods(http.MethodDelete)
	return router
}

func InitDB() *gorm.DB {
	// Build the PostgreSQL connection string
	PostgresURI := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		"localhost", "5432", "postgres", "stocksdb", "root")

	gormDB, err := gorm.Open(postgres.Open(PostgresURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect DB %v", err)
	}
	log.Println("DB Connection successful")

	return gormDB
}
func validateDB(gormDB *gorm.DB) {
	postgresDB, err := gormDB.DB() // Get the underlying sql.DB
	if err != nil {
		log.Fatalf("failed to fetch underlying sql.DB %v", err)
	}
	postgresDB.Ping() // Ping the database
	if err != nil {
		log.Fatalf("failed to ping DB %v", err)
	}
	log.Println("Database Connection ping successful...!")
}

func main() {
	fmt.Println("Hello, World!")

	// In memory store
	// repo := store.NewStockRepository()

	//DB conn
	db := InitDB()
	validateDB(db)

	// Postgres store
	repo := store.NewStockRepositoryDB(db)

	service := service.NewStockService(repo)
	handler := handlers.NewStockHandler(service)
	router := setupRouter(handler)
	ser := http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	// Channel for Interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Running on saparate goroutine
	go func() {
		log.Println("Server will start")
		err := ser.ListenAndServe()
		if err != nil {
			log.Fatalln("HTTP Server Error", err)
		}
	}()

	// Wait for Interrupt signal
	<-stop

	log.Println("Stocks API server stopped successfully")

}
