package providers

import (
	"api-loyalty-point-agent/businesses/providers"

	"time"

	"gorm.io/gorm"
)

type Provider struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name" gorm:"unique"`
	URL       string         `json:"url" gorm:"unique"`
}

func (rec *Provider) ToDomain() providers.Domain {
	return providers.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		Name:      rec.Name,
		URL:       rec.URL,
	}
}

func FromDomain(domain *providers.Domain) *Provider {
	return &Provider{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Name:      domain.Name,
		URL:       domain.URL,
	}
}
