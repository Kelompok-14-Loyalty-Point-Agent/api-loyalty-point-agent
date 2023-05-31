package response

import (
	"api-loyalty-point-agent/businesses/providers"
	"api-loyalty-point-agent/businesses/stocks"
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID         uint             `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
	DeletedAt  gorm.DeletedAt   `json:"deleted_at"`
	Name       string           `json:"name"`
	Type       string           `json:"type"`
	Stock      float64          `json:"stock"`
	Price      float64          `json:"price"`
	Quantity   float64          `json:"quantity"`
	ProviderID uint             `json:"provider_id"`
	Provider   providers.Domain `json:"-"`
}

func FromDomain(domain stocks.Domain) Stock {
	return Stock{
		ID:         domain.ID,
		Name:       domain.Name,
		Type:       domain.Type,
		Stock:      domain.Stock,
		Price:      domain.Price,
		Provider:   domain.Provider,
		ProviderID: domain.ProviderID,
		Quantity:   domain.Quantity,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:  domain.DeletedAt,
	}
}
