package transactions

import (
	"api-loyalty-point-agent/businesses/transactions"
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
	StockDetails   stock_details.StockDetail `json:"Stock_details" gorm:"foreignKey:stock_details_id"`
	StockID        uint                      `json:"stock_id"`
	Price          float64                   `json:"price"`
	Product        string                    `json:"product"`
	Payment_method string                    `json:"payment_method"`
	Point          uint                      `json:"point"`
	Status         string                    `json:"status" gorm:"type:enum('success', 'failed', 'on-process');default:'success';not_null"`
	Description    string                    `json:"description" gorm:"type:enum('top up');default:'top up';not_null"`
	UserID         uint                      `json:"user_id"`
	// TotalTransactionMade int64                     `json:"total_transaction_made"`
}

func (rec *Transaction) ToDomain() transactions.Domain {
	return transactions.Domain{
		ID:             rec.ID,
		CreatedAt:      rec.CreatedAt,
		UpdatedAt:      rec.UpdatedAt,
		DeletedAt:      rec.DeletedAt,
		Phone:          rec.Phone,
		StockDetailsID: rec.StockDetailsID,
		StockDetails:   rec.StockDetails.ToDomain(),
		StockID:        rec.StockID,
		Price:          rec.Price,
		Product:        rec.Product,
		Payment_method: rec.Payment_method,
		Point:          rec.Point,
		Status:         rec.Status,
		Description:    rec.Description,
		UserID:         rec.UserID,
	}
}

func FromDomain(domain *transactions.Domain) *Transaction {
	return &Transaction{
		ID:             domain.ID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		Phone:          domain.Phone,
		StockDetailsID: domain.StockDetailsID,
		StockID:        domain.StockID,
		Price:          domain.Price,
		Product:        domain.Product,
		Payment_method: domain.Payment_method,
		Point:          domain.Point,
		Status:         domain.Status,
		Description:    domain.Description,
		UserID:         domain.UserID,
	}
}
