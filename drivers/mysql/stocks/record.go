package stocks

import (
	"api-loyalty-point-agent/businesses/stocks"
	"api-loyalty-point-agent/drivers/mysql/providers"
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID         uint               `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	DeletedAt  gorm.DeletedAt     `json:"deleted_at" gorm:"index"`
	Name       string             `json:"name"`
	Type       string             `json:"type"`
	Stock      float64            `json:"stock"`
	Price      float64            `json:"price"`
	Quantity   float64            `json:"quantity"`
	ProviderID uint               `json:"provider_id"`
	Provider   providers.Provider `json:"-" gorm:"foreignKey:ProviderID"`
}

func (rec *Stock) ToDomain() stocks.Domain {
	return stocks.Domain{
		ID:         rec.ID,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
		DeletedAt:  rec.DeletedAt,
		Name:       rec.Name,
		Type:       rec.Type,
		Stock:      rec.Stock,
		Price:      rec.Price,
		Quantity:   rec.Quantity,
		Provider:   rec.ToDomain().Provider,
		ProviderID: rec.ProviderID,
	}
}

func FromDomain(domain *stocks.Domain) *Stock {
	return &Stock{
		ID:         domain.ID,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:  domain.DeletedAt,
		Name:       domain.Name,
		Type:       domain.Type,
		Stock:      domain.Stock,
		Price:      domain.Price,
		Quantity:   domain.Quantity,
		Provider:   *providers.FromDomain(&domain.Provider),
		ProviderID: domain.ProviderID,
	}
}
