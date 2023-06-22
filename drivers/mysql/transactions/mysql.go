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
		transactions = append(transactions, transaction.ToDomain())
	}

	return transactions, nil
}

func (cr *transactionRepository) GetByID(ctx context.Context, id string) (transactions.Domain, error) {
	var transaction Transaction

	if err := cr.conn.WithContext(ctx).Preload("StockDetails").First(&transaction, "id = ?", id).Error; err != nil {
		return transactions.Domain{}, err
	}

	return transaction.ToDomain(), nil
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

	// product name
	record.Product += strings.Title(stock.Type) + " " + provider.Name

	// point
	record.Point = uint(stock_detail.Price) / 1000

	if err := cr.conn.WithContext(ctx).Save(&stock_detail).Error; err != nil {
		return transactions.Domain{}, err
	}

	if err := cr.conn.WithContext(ctx).Save(&stock).Error; err != nil {
		return transactions.Domain{}, err
	}

	if err := cr.conn.WithContext(ctx).Save(&record).Error; err != nil {
		return transactions.Domain{}, err
	}

	// Mengambil data profil berdasarkan record.UserID
	var profile profiles.Profile
	if err := cr.conn.WithContext(ctx).First(&profile, "id = ?", record.UserID).Error; err != nil {
		return transactions.Domain{}, err
	}

	// Increment total transaksi pada profil
	profile.TransactionMade += 1

	// // Dapatkan bulan dan tahun saat ini
	// currentMonth := time.Now().Month()
	// currentYear := time.Now().Year()

	// // Lakukan query untuk menghitung total transaksi bulan ini
	// var totalTransactionsThisMonth int64
	// if err := cr.conn.WithContext(ctx).Model(&transactions.Domain{}).
	// 	Joins("JOIN profiles ON transactions.user_id = profiles.id").
	// 	Where("profiles.id = ? AND MONTH(transactions.created_at) = ? AND YEAR(transactions.created_at) = ?", record.UserID, currentMonth, currentYear).
	// 	Count(&totalTransactionsThisMonth).Error; err != nil {
	// 	return transactions.Domain{}, err
	// }

	// // Perbarui nilai TotalTransactionsThisMonth pada profil
	// profile.MonthlyTransaction = uint(totalTransactionsThisMonth)

	// Simpan perubahan pada profil
	if err := cr.conn.WithContext(ctx).Save(&profile).Error; err != nil {
		return transactions.Domain{}, err
	}

	var records []Transaction

	if err := cr.conn.WithContext(ctx).Preload("StockDetails").
		Where("user_id = ? AND MONTH(created_at) = ?", record.UserID, time.Now().Month()).
		Find(&records).Error; err != nil {
		return transactions.Domain{}, err
	}

	profile.MonthlyTransaction = uint(len(records))

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

	if err := cr.conn.WithContext(ctx).Save(&updatedTransaction).Error; err != nil {
		return transactions.Domain{}, err
	}

	return updatedTransaction.ToDomain(), nil
}

// func (cr *transactionRepository) GetTotalTransactionMade(ctx context.Context, userid string) (transactions.Domain, error) {
// 	var count int64
// 	err := cr.conn.WithContext(ctx).Model(&Transaction{}).Where("user_id = ?", userid).Count(&count).Error
// 	for _, transaction := range records {
// 		transactions = append(transactions, transaction.ToDomain())
// 	}
// 	if err != nil {
// 		return 0, err
// 	}
// 	return transactions, nil
// }
