package store

import (
	"errors"
	"log"
	"stocksapi/models"

	"gorm.io/gorm"
)

type StockRepositoryDBImpl struct {
	DBConn *gorm.DB
}

func NewStockRepositoryDB(db *gorm.DB) models.StockRepository {
	return &StockRepositoryDBImpl{
		DBConn: db,
	}
}

func (r *StockRepositoryDBImpl) GetStocks() ([]models.Stock, error) {

	var allstocks []models.Stock
	result := r.DBConn.Raw("SELECT * FROM stocks").Scan(&allstocks)
	if result.Error != nil {
		log.Println("GetStocks:DB Query failed", result.Error)
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		log.Println("GetStocks: No stocks found")
		return nil, errors.New("no stocks found")
	}
	log.Println("CreateStock:DB Query Successful for all stocks ", allstocks)
	return allstocks, nil
}

func (r *StockRepositoryDBImpl) GetStockByID(id int) (*models.Stock, error) {
	var stock models.Stock
	result := r.DBConn.Raw("SELECT * FROM stocks WHERE id = ?", id).Scan(&stock)
	if result.Error != nil {
		log.Println("GetStockByID:DB Query failed", result.Error)
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		log.Println("GetStockByID: No stock found with id", id)
		return nil, errors.New("no stock found")
	}
	log.Println("CreateStock:DB Query Successful for stock id : ", stock.Id)
	return &stock, nil

}

func (r *StockRepositoryDBImpl) CreateStock(stock *models.Stock) error {
	result := r.DBConn.Exec("INSERT INTO stocks (id, name, company, price) VALUES (?, ?, ?, ?)", stock.Id, stock.Name, stock.Company, stock.Price)
	if result.Error != nil {
		log.Println("CreateStock:DB Insertion failed", result.Error)
		return result.Error
	}
	log.Println("CreateStock:DB Insertion Successful for stock id : ", stock.Id)
	return nil
}

func (r *StockRepositoryDBImpl) UpdateStock(stock *models.Stock) error {
	result := r.DBConn.Exec("UPDATE stocks SET name = ?, company = ?, price = ? where id = ?", stock.Name, stock.Company, stock.Price, stock.Id)
	if result.Error != nil {
		log.Println("UpdateStock:DB Update failed", result.Error)
		return result.Error
	}
	log.Println("UpdateStock:DB Update Successful for stock id : ", stock.Id)
	return nil
}

func (r *StockRepositoryDBImpl) DeleteStock(id int) error {
	result := r.DBConn.Exec("DELETE FROM stocks where id = ?", id)
	if result.Error != nil {
		log.Println("DeleteStock:DB Delete failed", result.Error)
		return result.Error
	}
	log.Println("DeleteStock:DB Delete Successful for stock id : ", id)
	return nil
}
