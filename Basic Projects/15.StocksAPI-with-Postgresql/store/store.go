package store

import (
	"errors"
	"stocksapi/models"
	"sync"
)

type StockRepositoryImpl struct {
	sync.Mutex
	stocks map[int]*models.Stock
}

func NewStockRepository() models.StockRepository {
	return &StockRepositoryImpl{
		stocks: make(map[int]*models.Stock),
	}
}

func (r *StockRepositoryImpl) GetStocks() ([]models.Stock, error) {
	r.Lock()
	defer r.Unlock()
	var stocks []models.Stock
	for _, each := range r.stocks {
		stocks = append(stocks, *each)
	}
	return stocks, nil
}

func (r *StockRepositoryImpl) GetStockByID(id int) (*models.Stock, error) {
	r.Lock()
	defer r.Unlock()
	stock, exist := r.stocks[id]
	if !exist {
		return stock, errors.New("stock not found")
	}
	return stock, nil
}

func (r *StockRepositoryImpl) CreateStock(stock *models.Stock) error {
	r.Lock()
	defer r.Unlock()
	_, exist := r.stocks[stock.Id]
	if exist {
		return errors.New("stock is already exist")
	}
	r.stocks[stock.Id] = stock
	return nil
}

func (r *StockRepositoryImpl) UpdateStock(stock *models.Stock) error {
	r.Lock()
	defer r.Unlock()
	_, exist := r.stocks[stock.Id]
	if !exist {
		return errors.New("stock not found")
	}
	r.stocks[stock.Id] = stock
	return nil
}

func (r *StockRepositoryImpl) DeleteStock(id int) error {
	r.Lock()
	defer r.Unlock()
	_, exist := r.stocks[id]
	if !exist {
		return errors.New("stock not found")
	}
	delete(r.stocks, id)
	return nil
}
