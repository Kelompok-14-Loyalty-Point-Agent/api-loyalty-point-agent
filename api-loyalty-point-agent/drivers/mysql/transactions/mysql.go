package transactions

import (
	"api-loyalty-point-agent/businesses/transactions"
	"context"

	"gorm.io/gorm"
)

type transactionRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) transactions.Repository {
	return &transactionRepository{
		conn: conn,
	}
}

func (cr *transactionRepository) GetAll(ctx context.Context) ([]transactions.Domain, error) {
	var records []Transaction

	if err := cr.conn.WithContext(ctx).Preload("StockDetails").Find(&records).Error; err != nil {
		return nil, err
	}

	transactions := []transactions.Domain{}

	// fmt.Println(transactions)

	for _, transaction := range records {
		transactions = append(transactions, transaction.ToDomain())
	}

	return transactions, nil
}

func (cr *transactionRepository) GetByID(ctx context.Context, id string) (transactions.Domain, error) {
	var transaction Transaction

	if err := cr.conn.WithContext(ctx).First(&transaction, "id = ?", id).Error; err != nil {
		return transactions.Domain{}, err
	}

	return transaction.ToDomain(), nil
}

func (cr *transactionRepository) Create(ctx context.Context, transactionDomain *transactions.Domain) (transactions.Domain, error) {
	record := FromDomain(transactionDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	//---------------------------

	if err := result.Error; err != nil {
		return transactions.Domain{}, err
	}

	if err := result.Preload("StockDetails").Last(&record).Error; err != nil {
		return transactions.Domain{}, err
	}

	return record.ToDomain(), nil
}
