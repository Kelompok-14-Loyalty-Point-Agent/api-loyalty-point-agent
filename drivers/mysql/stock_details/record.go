package stock_details

import (
	"api-loyalty-point-agent/businesses/stock_details"

	"time"

	"gorm.io/gorm"
)

type StockDetail struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Stock     float64        `json:"stock"`
	Price     float64        `json:"price"`
	Quantity  float64        `json:"quantity"`
	StockID   uint           `json:"stock_id"`
}

func (rec *StockDetail) ToDomain() stock_details.Domain {
	return stock_details.Domain{
		ID:         rec.ID,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
		DeletedAt:  rec.DeletedAt,
		Stock:      rec.Stock,
		Price:      rec.Price,
		Quantity:   rec.Quantity,
		StockID:    rec.StockID,
	}
}

func FromDomain(domain *stock_details.Domain) *StockDetail {
	return &StockDetail{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Stock:     domain.Stock,
		Price:     domain.Price,
		Quantity:  domain.Quantity,
		StockID:   domain.StockID,
	}
}
