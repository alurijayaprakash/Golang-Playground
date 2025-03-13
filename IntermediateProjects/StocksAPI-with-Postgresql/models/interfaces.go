package models

import "net/http"

type Stock struct {
	Id      int
	Name    string
	Company string
	Price   int
}

// Store
type StockRepository interface {
	GetStocks() ([]Stock, error)         // Get all stocks
	GetStockByID(id int) (*Stock, error) // Get a stock by its ID.
	CreateStock(stock *Stock) error      // Create a new stock
	UpdateStock(stock *Stock) error      // Update an existing stock
	DeleteStock(id int) error            // Delete a stock by its ID
}

// Service
type StockService interface {
	ListStocks() ([]Stock, error)    // List all stocks
	GetStock(id int) (*Stock, error) // Get a stock by ID
	AddStock(stock *Stock) error     // Add a new stock
	ModifyStock(stock *Stock) error  // Modify an existing stock
	RemoveStock(id int) error        // Remove a stock by ID
}

type StockHandler interface {
	ListStocksHandler(w http.ResponseWriter, r *http.Request)  // List all stocks
	GetStockHandler(w http.ResponseWriter, r *http.Request)    // Get a stock by ID
	CreateStockHandler(w http.ResponseWriter, r *http.Request) // Create a new stock
	UpdateStockHandler(w http.ResponseWriter, r *http.Request) // Update an existing stock
	DeleteStockHandler(w http.ResponseWriter, r *http.Request) // Delete a stock by ID

}
