package response

import (
	"api-loyalty-point-agent/businesses/stock_details"
	"time"

	"gorm.io/gorm"
)

type StockDetail struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	StockID   uint           `json:"stock_id"`
	Stock     float64        `json:"stock"`
	Price     float64        `json:"price"`
	Quantity  float64        `json:"quantity"`
}

func FromDomain(domain stock_details.Domain) StockDetail {
	return StockDetail{
		ID:        domain.ID,
		StockID:   domain.StockID,
		Stock:     domain.Stock,
		Price:     domain.Price,
		Quantity:  domain.Quantity,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
