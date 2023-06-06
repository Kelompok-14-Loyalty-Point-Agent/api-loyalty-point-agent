package stock_transactions

import (
	"api-loyalty-point-agent/businesses/stock_transactions"
	"context"

	"gorm.io/gorm"
)

type stock_transactionRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) stock_transactions.Repository {
	return &stock_transactionRepository{
		conn: conn,
	}
}

func (cr *stock_transactionRepository) GetAll(ctx context.Context) ([]stock_transactions.Domain, error) {
	var records []StockTransaction

	if err := cr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	stock_transactions := []stock_transactions.Domain{}

	for _, stock_transaction := range records {
		stock_transactions = append(stock_transactions, stock_transaction.ToDomain())
	}

	return stock_transactions, nil
}

func (cr *stock_transactionRepository) GetByID(ctx context.Context, id string) (stock_transactions.Domain, error) {
	var stock_transaction StockTransaction

	if err := cr.conn.WithContext(ctx).First(&stock_transaction, "id = ?", id).Error; err != nil {
		return stock_transactions.Domain{}, err
	}

	return stock_transaction.ToDomain(), nil
}
