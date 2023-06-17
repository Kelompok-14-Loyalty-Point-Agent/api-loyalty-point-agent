package response

import (
	"api-loyalty-point-agent/businesses/providers"
	"time"

	"gorm.io/gorm"
)

type Provider struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name      string         `json:"name"`
	URL       string         `json:"url"`
}

func FromDomain(domain providers.Domain) Provider {
	return Provider{
		ID:        domain.ID,
		Name:      domain.Name,
		URL:       domain.URL,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
