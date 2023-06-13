package stocks

import (
	"api-loyalty-point-agent/businesses/stock_transactions"
	"api-loyalty-point-agent/businesses/stocks"
	_dbStockTransaction "api-loyalty-point-agent/drivers/mysql/stock_transactions"
	"context"
	"errors"

	"gorm.io/gorm"
)

type stockRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) stocks.Repository {
	return &stockRepository{
		conn: conn,
	}
}

func (cr *stockRepository) GetAll(ctx context.Context) ([]stocks.Domain, error) {
	var records []Stock

	if err := cr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	stocks := []stocks.Domain{}

	for _, stock := range records {
		stocks = append(stocks, stock.ToDomain())
	}

	return stocks, nil
}

func (cr *stockRepository) GetByID(ctx context.Context, id string) (stocks.Domain, error) {
	var stock Stock

	if err := cr.conn.WithContext(ctx).First(&stock, "id = ?", id).Error; err != nil {
		return stocks.Domain{}, err
	}

	return stock.ToDomain(), nil
}

func (cr *stockRepository) AddStock(ctx context.Context, stock_transactionDomain *stock_transactions.Domain) (stock_transactions.Domain, error) {
	record := _dbStockTransaction.FromDomain(stock_transactionDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return stock_transactions.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return stock_transactions.Domain{}, err
	}

	var stock Stock

	if err := cr.conn.WithContext(ctx).First(&stock, "id = ?", record.StockID).Error; err != nil {
		return stock_transactions.Domain{}, err
	}

	if stock.Type == "data" {
		// 1 GB / Rp.10000
		record.PayAmount = record.InputStock * 10000
	} else if stock.Type == "credit" {
		if record.InputStock < 10000 {
			return stock_transactions.Domain{}, errors.New("credit input minimum is 10000")
		} else {
			// 10000 / Rp.12000
			record.PayAmount = record.InputStock / 10000 * 12000
		}

	}

	record.Status = "success"

	if err := cr.conn.WithContext(ctx).Save(&record).Error; err != nil {
		return stock_transactions.Domain{}, err
	}

	stock.TotalStock += record.InputStock
	stock.LastTopUp = record.CreatedAt

	if err := cr.conn.WithContext(ctx).Save(&stock).Error; err != nil {
		return stock_transactions.Domain{}, err
	}

	return record.ToDomain(), nil
}
