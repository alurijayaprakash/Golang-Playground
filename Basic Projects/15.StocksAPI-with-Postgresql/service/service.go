package service

import "stocksapi/models"

type StockServiceImpl struct {
	repo models.StockRepository
}

func NewStockService(repo models.StockRepository) models.StockService {
	return &StockServiceImpl{
		repo: repo,
	}
}

func (s *StockServiceImpl) ListStocks() ([]models.Stock, error) {
	return s.repo.GetStocks()
}

func (s *StockServiceImpl) GetStock(id int) (*models.Stock, error) {
	return s.repo.GetStockByID(id)
}

func (s *StockServiceImpl) AddStock(stock *models.Stock) error {
	return s.repo.CreateStock(stock)
}

func (s *StockServiceImpl) ModifyStock(stock *models.Stock) error {
	return s.repo.UpdateStock(stock)
}

func (s *StockServiceImpl) RemoveStock(id int) error {
	return s.repo.DeleteStock(id)
}
