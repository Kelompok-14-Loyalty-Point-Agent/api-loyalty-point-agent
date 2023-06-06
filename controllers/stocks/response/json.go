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
	Type       string           `json:"type"`
	TotalStock float64          `json:"stock"`
	LastTopUp  time.Time        `json:"last_top_up"`
	ProviderID uint             `json:"provider_id"`
	Provider   providers.Domain `json:"-"`
}

func FromDomain(domain stocks.Domain) Stock {
	return Stock{
		ID:         domain.ID,
		Type:       domain.Type,
		TotalStock: domain.TotalStock,
		LastTopUp:  domain.LastTopUp,
		Provider:   domain.Provider,
		ProviderID: domain.ProviderID,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:  domain.DeletedAt,
	}
}
