package response

import (
	"api-loyalty-point-agent/businesses/stock_transactions"
	"time"

	"gorm.io/gorm"
)

type StockTransaction struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	ProviderName  string         `json:"provider_name"`
	InputStock    float64        `json:"input_stock"`
	PayAmount     float64        `json:"pay_amount" `
	PaymentMethod string         `json:"payment_method" `
	Status        string         `json:"status" `
	StockID       uint           `json:"stock_id"`
	Stocks        string         `json:"-"`
	UserID        uint           `json:"-"`
}

func FromDomain(domain stock_transactions.Domain) StockTransaction {
	return StockTransaction{
		ID:            domain.ID,
		ProviderName:  domain.ProviderName,
		InputStock:    domain.InputStock,
		PayAmount:     domain.PayAmount,
		PaymentMethod: domain.PaymentMethod,
		Status:        domain.Status,
		StockID:       domain.StockID,
		Stocks:        domain.Stocks,
		UserID:        domain.UserID,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
		DeletedAt:     domain.DeletedAt,
	}
}
