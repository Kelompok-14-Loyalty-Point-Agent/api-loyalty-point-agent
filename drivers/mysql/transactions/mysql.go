package transactions

import (
	"api-loyalty-point-agent/businesses/transactions"
	"errors"
	"strings"
	"time"

	"api-loyalty-point-agent/drivers/mysql/profiles"
	"api-loyalty-point-agent/drivers/mysql/providers"
	"api-loyalty-point-agent/drivers/mysql/stock_details"
	"api-loyalty-point-agent/drivers/mysql/stocks"
	"api-loyalty-point-agent/utils"
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

	for _, transaction := range records {

		var profile profiles.Profile
		_ = cr.conn.WithContext(ctx).First(&profile, "id = ?", transaction.UserID).Error
		
		domainTransaction := transaction.ToDomain()
		domainTransaction.UserName = profile.Name
		domainTransaction.URL = profile.URL
		domainTransaction.Member = profile.Member

		if len(domainTransaction.UserName) != 0 {
			transactions = append(transactions, domainTransaction)
		}
	}

	return transactions, nil
}

func (cr *transactionRepository) GetByID(ctx context.Context, id string) (transactions.Domain, error) {
	var transaction Transaction

	if err := cr.conn.WithContext(ctx).Preload("StockDetails").First(&transaction, "id = ?", id).Error; err != nil {
		return transactions.Domain{}, err
	}

	var profile profiles.Profile
		if err := cr.conn.WithContext(ctx).First(&profile, "id = ?", transaction.UserID).Error; err != nil {
			return transactions.Domain{}, err
	}

	domainTransaction := transaction.ToDomain()
	domainTransaction.UserName = profile.Name
	domainTransaction.URL = profile.URL
	domainTransaction.Member = profile.Member

	return domainTransaction, nil
}

func (cr *transactionRepository) Create(ctx context.Context, transactionDomain *transactions.Domain) (transactions.Domain, error) {
	record := FromDomain(transactionDomain)

	var stock_detail stock_details.StockDetail

	if err := cr.conn.WithContext(ctx).First(&stock_detail, "id = ?", transactionDomain.StockDetailsID).Error; err != nil {
		return transactions.Domain{}, err
	}

	if stock_detail.Quantity-1 < 0 {
		err := errors.New("run out of quantity")
		return transactions.Domain{}, err
	}

	// quantity
	stock_detail.Quantity -= 1

	var stock stocks.Stock

	if err := cr.conn.WithContext(ctx).First(&stock, "id = ?", stock_detail.StockID).Error; err != nil {
		return transactions.Domain{}, err
	}

	if stock.TotalStock-stock_detail.Stock < 0 {
		err := errors.New("run out of stock")
		return transactions.Domain{}, err
	}

	// total stock
	stock.TotalStock -= stock_detail.Stock

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return transactions.Domain{}, err
	}

	if err := result.Preload("StockDetails").Last(&record).Error; err != nil {
		return transactions.Domain{}, err
	}

	var provider providers.Provider

	if err := cr.conn.WithContext(ctx).First(&provider, "id = ?", stock.ProviderID).Error; err != nil {
		return transactions.Domain{}, err
	}

	record.StockID = stock_detail.StockID
	record.Price = stock_detail.Price

	// product name
	record.Product += strings.Title(stock.Type) + " " + provider.Name

	// point
	record.Point = float32(stock_detail.Price) / 1000


	var profile profiles.Profile
	if err := cr.conn.WithContext(ctx).First(&profile, "id = ?", record.UserID).Error; err != nil {
		return transactions.Domain{}, err
	}

	// transaction_made
	profile.TransactionMade += 1

	// point to profile
	profile.Point += record.Point
	utils.RoundFloat(float64(profile.Point), 1)

	// monthly transaction
	var records []Transaction

	if err := cr.conn.WithContext(ctx).Preload("StockDetails").
		Where("user_id = ? AND MONTH(created_at) = ?", record.UserID, time.Now().Month()).
		Find(&records).Error; err != nil {
		return transactions.Domain{}, err
	}

	profile.MonthlyTransaction = uint(len(records))

	// member
	if profile.TransactionMade >= 5 && profile.TransactionMade <= 10 {
		profile.Member = "silver"
	} else if profile.TransactionMade >= 11 && profile.TransactionMade <= 15 {
		profile.Member = "gold"
	} else if profile.TransactionMade >= 16 {
		profile.Member = "platinum"
	}

	if err := cr.conn.WithContext(ctx).Save(&profile).Error; err != nil {
		return transactions.Domain{}, err
	}

	if err := cr.conn.WithContext(ctx).Save(&stock_detail).Error; err != nil {
		return transactions.Domain{}, err
	}

	if err := cr.conn.WithContext(ctx).Save(&stock).Error; err != nil {
		return transactions.Domain{}, err
	}

	if err := cr.conn.WithContext(ctx).Save(&record).Error; err != nil {
		return transactions.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (cr *transactionRepository) GetAllByUserID(ctx context.Context, userid string) ([]transactions.Domain, error) {
	var records []Transaction

	if err := cr.conn.WithContext(ctx).Preload("StockDetails").Find(&records, `user_id = ?`, userid).Error; err != nil {
		return nil, err
	}

	transactions := []transactions.Domain{}

	for _, transaction := range records {
		transactions = append(transactions, transaction.ToDomain())
	}

	return transactions, nil
}

func (cr *transactionRepository) GetAllByUserIDSorted(ctx context.Context, userid string) ([]transactions.Domain, error) {
	var records []Transaction

	if err := cr.conn.WithContext(ctx).Preload("StockDetails").Order("created_at DESC").Find(&records, `user_id = ?`, userid).Error; err != nil {
		return nil, err
	}

	transactions := []transactions.Domain{}

	for _, transaction := range records {
		transactions = append(transactions, transaction.ToDomain())
	}

	return transactions, nil
}

func (cr *transactionRepository) UpdatePoint(ctx context.Context, transactionDomain *transactions.Domain, id string) (transactions.Domain, error) {
	transaction, err := cr.GetByID(ctx, id)

	if err != nil {
		return transactions.Domain{}, err
	}

	updatedTransaction := FromDomain(&transaction)

	if transactionDomain.Point < 0 {
		err := errors.New("invalid point")
		return transactions.Domain{}, err
	}

	updatedTransaction.Point = transactionDomain.Point
	updatedTransaction.StockDetailsID = transactionDomain.StockDetailsID
	updatedTransaction.StockDetails = stock_details.StockDetail(transaction.StockDetails)
	updatedTransaction.StockDetails.Price = transaction.Price

	var profile profiles.Profile
	if err := cr.conn.WithContext(ctx).First(&profile, "id = ?", transaction.UserID).Error; err != nil {
		return transactions.Domain{}, err
	}

	if (profile.Point - transaction.Point) + transactionDomain.Point < 0 {
		return transactions.Domain{}, errors.New("user already spent their points, updating their point is not valid right now")
	}else{
		profile.Point = (profile.Point - transaction.Point) + transactionDomain.Point
	}
	
	if err := cr.conn.WithContext(ctx).Save(&profile).Error; err != nil {
		return transactions.Domain{}, err
	}

	if err := cr.conn.WithContext(ctx).Save(&updatedTransaction).Error; err != nil {
		return transactions.Domain{}, err
	}

	return updatedTransaction.ToDomain(), nil
}
