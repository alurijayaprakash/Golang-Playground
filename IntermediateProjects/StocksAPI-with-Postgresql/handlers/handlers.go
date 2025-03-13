package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"stocksapi/models"
	"strconv"

	"github.com/gorilla/mux"
)

type StockHandlerImpl struct {
	Service models.StockService
}

func NewStockHandler(ser models.StockService) models.StockHandler {
	return &StockHandlerImpl{
		Service: ser,
	}
}

func (h *StockHandlerImpl) ListStocksHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("ListStocksHandler...!")
	w.Header().Set("Content-Type", "application/json")

	getStocks, err := h.Service.ListStocks()
	if err != nil {
		log.Println("ListStocksHandler:Unable to fetch the stocks", err)
		respondWithError(w, http.StatusInternalServerError, "Unable to fetch stocks")
		return
	}
	log.Println("getStocks", getStocks)
	err = json.NewEncoder(w).Encode(getStocks)
	if err != nil {
		log.Println("ListStocksHandler:Unable to encode the stocks", err)
		respondWithError(w, http.StatusInternalServerError, "Unable to encode the stocks")
		return
	}
}

func (h *StockHandlerImpl) GetStockHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetStockHandler...!")
	w.Header().Set("Content-Type", "application/json")
	var stock *models.Stock
	params := mux.Vars(r)
	stockid, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("GetStockHandler:Unable to read client data", err)
		respondWithError(w, http.StatusBadRequest, "Unable to read client data")
		return
	}

	stock, err = h.Service.GetStock(stockid)
	if err != nil {
		log.Println("GetStockHandler:stock data not found", err)
		respondWithError(w, http.StatusNotFound, "stock data not found")
		return
	}
	err = json.NewEncoder(w).Encode(stock)
	if err != nil {
		log.Println("GetStockHandler:encode error", err)
		respondWithError(w, http.StatusNotFound, "encode error")
		return
	}

}
func (h *StockHandlerImpl) CreateStockHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("CreateStockHandler...!")
	w.Header().Set("Content-Type", "application/json")
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Println("CreateStockHandler:unable to decode client body", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Create a new stock
	err = h.Service.AddStock(&stock)
	if err != nil {
		log.Println("CreateStockHandler:error while creating stock", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithSuccess(w, http.StatusCreated, "stock addedd successfully")
}
func (h *StockHandlerImpl) UpdateStockHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("UpdateStockHandler...!")
	w.Header().Set("Content-Type", "application/json")
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Println("UpdateStockHandler:unable to decode client body", err)
		respondWithError(w, http.StatusBadRequest, "unable to read client body")
		return
	}

	_, err = h.Service.GetStock(stock.Id)
	if err != nil {
		log.Println("UpdateStockHandler:stock not found", err)
		respondWithError(w, http.StatusNotFound, "stock not found")
		return
	}

	// Create a new stock
	err = h.Service.ModifyStock(&stock)
	if err != nil {
		log.Println("UpdateStockHandler:error while modifying stock", err)
		respondWithError(w, http.StatusInternalServerError, "error while modifying stock")
		return
	}
	respondWithSuccess(w, http.StatusOK, "stock updated successfully")
}
func (h *StockHandlerImpl) DeleteStockHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("DeleteStockHandler...!")
	w.Header().Set("Content-Type", "application/json")
	// var stock *models.Stock
	params := mux.Vars(r)
	stockid, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("DeleteStockHandler:Unable to read client data", err)
		respondWithError(w, http.StatusBadRequest, "Unable to read client data")
		return
	}

	_, err = h.Service.GetStock(stockid)
	if err != nil {
		log.Println("DeleteStockHandler:stock data not found", err)
		respondWithError(w, http.StatusNotFound, "stock data not found")
		return
	}
	err = h.Service.RemoveStock(stockid)
	if err != nil {
		log.Println("DeleteStockHandler:unable to delete stock", err)
		respondWithError(w, http.StatusInternalServerError, "unable to delete stock")
		return
	}
	respondWithSuccess(w, http.StatusOK, "stock deleted successfully")
}

func respondWithError(w http.ResponseWriter, statuscode int, error_msg string) {
	w.WriteHeader(statuscode)
	err := json.NewEncoder(w).Encode(map[string]string{
		"error": error_msg,
	})
	if err != nil {
		log.Printf("respondWithError:Unable to encode error message")
	}
}

func respondWithSuccess(w http.ResponseWriter, statuscode int, msg string) {
	w.WriteHeader(statuscode)
	err := json.NewEncoder(w).Encode(map[string]string{
		"status": msg,
	})
	if err != nil {
		log.Printf("respondWithSuccess:Unable to encode error message")
	}
}
