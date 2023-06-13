package response

import (
	"api-loyalty-point-agent/businesses/stocks"
	"api-loyalty-point-agent/businesses/transaction"

	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	// Amount         uint           `json:"amount"`
	Phone          string        `json:"phone"`
	StockID        uint          `json:"stock_id"`
	Stock          stocks.Domain `json:"stock"`
	Type           string        `json:"type"`
	Product        uint          `json:"product"`
	Payment_method uint          `json:"payment_method"`
	Point          uint          `json:"point"`
}

func FromDomain(domain transaction.Domain) Transaction {
	return Transaction{
		ID: domain.ID,
		// Amount:         domain.Amount,
		Phone:          domain.Phone,
		Product:        domain.Product,
		StockID:        domain.StockID,
		Stock:          domain.Stock,
		Type:           domain.Stock.Type,
		Payment_method: domain.Payment_method,
		Point:          domain.Point,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
	}
}
