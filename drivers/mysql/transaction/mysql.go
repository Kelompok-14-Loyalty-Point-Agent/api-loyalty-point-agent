package transaction

import (
	"api-loyalty-point-agent/businesses/transaction"
	"context"

	"gorm.io/gorm"
)

type transactionRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) transaction.Repository {
	return &transactionRepository{
		conn: conn,
	}
}

func (cr *transactionRepository) GetAll(ctx context.Context) ([]transaction.Domain, error) {
	var records []Transaction

	if err := cr.conn.WithContext(ctx).Preload("StockDetails").Find(&records).Error; err != nil {
		return nil, err
	}

	transactions := []transaction.Domain{}

	// fmt.Println(transactions)

	for _, transaction := range records {
		transactions = append(transactions, transaction.ToDomain())
	}

	return transactions, nil
}

func (cr *transactionRepository) GetByID(ctx context.Context, id string) (transaction.Domain, error) {
	var transactions Transaction

	if err := cr.conn.WithContext(ctx).First(&transactions, "id = ?", id).Error; err != nil {
		return transaction.Domain{}, err
	}

	return transactions.ToDomain(), nil
}

func (cr *transactionRepository) Create(ctx context.Context, transactionDomain *transaction.Domain) (transaction.Domain, error) {
	record := FromDomain(transactionDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	//---------------------------

	if err := result.Error; err != nil {
		return transaction.Domain{}, err
	}

	if err := result.Preload("StockDetails").Last(&record).Error; err != nil {
		return transaction.Domain{}, err
	}

	return record.ToDomain(), nil
}
