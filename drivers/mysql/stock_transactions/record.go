package stock_transactions

import (
	"api-loyalty-point-agent/businesses/stock_transactions"
	"time"

	"gorm.io/gorm"
)

type StockTransaction struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ProviderName  string         `json:"provider_name"`
	InputStock    float64        `json:"input_stock"`
	PayAmount     float64        `json:"pay_amount"`
	PaymentMethod string         `json:"payment_method"`
	Status        string         `json:"status" gorm:"type:enum('success', 'failed');default:'failed';not_null"`
	StockID       uint           `json:"stock_id"`
	UserID        uint           `json:"user_id"`
}

func (rec *StockTransaction) ToDomain() stock_transactions.Domain {
	return stock_transactions.Domain{
		ID:            rec.ID,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
		DeletedAt:     rec.DeletedAt,
		ProviderName:  rec.ProviderName,
		InputStock:    rec.InputStock,
		PayAmount:     rec.PayAmount,
		PaymentMethod: rec.PaymentMethod,
		Status:        rec.Status,
		StockID:       rec.StockID,
		UserID:        rec.UserID,
	}
}

func FromDomain(domain *stock_transactions.Domain) *StockTransaction {
	return &StockTransaction{
		ID:            domain.ID,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
		DeletedAt:     domain.DeletedAt,
		ProviderName:  domain.ProviderName,
		InputStock:    domain.InputStock,
		PayAmount:     domain.PayAmount,
		PaymentMethod: domain.PaymentMethod,
		Status:        domain.Status,
		StockID:       domain.StockID,
		UserID:        domain.UserID,
	}
}
