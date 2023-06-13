package transaction

import (
	"api-loyalty-point-agent/businesses/transaction"
	"api-loyalty-point-agent/drivers/mysql/stock_details"

	// "api-loyalty-point-agent/drivers/mysql/stocks"

	// "api-loyalty-point-agent/drivers/mysql/stocks"

	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID             uint                      `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time                 `json:"created_at"`
	UpdatedAt      time.Time                 `json:"updated_at"`
	DeletedAt      gorm.DeletedAt            `json:"deleted_at" gorm:"index"`
	Phone          string                    `json:"phone"`
	StockDetailsID uint                      `json:"stock_details_id"`
	StockDetails   stock_details.StockDetail `json:"StockDetails" gorm:"foreignKey:stock_details_id"`
	Product        uint                      `json:"product"`
	Payment_method uint                      `json:"payment_method"`
	Point          uint                      `json:"point"`
}

func (rec *Transaction) ToDomain() transaction.Domain {
	return transaction.Domain{
		ID:             rec.ID,
		CreatedAt:      rec.CreatedAt,
		UpdatedAt:      rec.UpdatedAt,
		DeletedAt:      rec.DeletedAt,
		Phone:          rec.Phone,
		StockDetailsID: rec.StockDetailsID,
		StockDetails:   rec.StockDetails.ToDomain(),
		StockID:        rec.StockDetails.StockID,
		Price:          rec.StockDetails.Price,
		Product:        rec.Product,
		Payment_method: rec.Payment_method,
		Point:          rec.Point,
	}
}

func FromDomain(domain *transaction.Domain) *Transaction {
	return &Transaction{
		ID:             domain.ID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		Phone:          domain.Phone,
		StockDetailsID: domain.StockDetailsID,
		Product:        domain.Product,
		Payment_method: domain.Payment_method,
		Point:          domain.Point,
	}
}
