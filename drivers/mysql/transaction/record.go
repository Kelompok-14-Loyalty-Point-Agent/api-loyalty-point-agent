package transaction

import (
	"api-loyalty-point-agent/businesses/transaction"
	"api-loyalty-point-agent/drivers/mysql/stocks"

	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Phone          string         `json:"phone"`
	StockID        uint           `json:"stock_id"`
	Stock          stocks.Stock   `json:"-" gorm:"foreignKey:StockID"`
	Product        uint           `json:"product"`
	Payment_method uint           `json:"payment_method"`
	Point          uint           `json:"point"`
}

func (rec *Transaction) ToDomain() transaction.Domain {
	return transaction.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		// Amount:         rec.Amount,
		Phone:          rec.Phone,
		StockID:        rec.StockID,
		Stock:          rec.Stock.ToDomain(),
		Product:        rec.Product,
		Payment_method: rec.Payment_method,
		Point:          rec.Point,
	}
}

func FromDomain(domain *transaction.Domain) *Transaction {
	return &Transaction{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		// Amount:         domain.Amount,
		Phone:          domain.Phone,
		StockID:        domain.StockID,
		Stock:          *stocks.FromDomain(&domain.Stock),
		Product:        domain.Product,
		Payment_method: domain.Payment_method,
		Point:          domain.Point,
	}
}
