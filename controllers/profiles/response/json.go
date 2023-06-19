package response

import (
	"api-loyalty-point-agent/businesses/profiles"
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name      string         `json:"name" validate:"required"`
	Address   string         `json:"address" validate:"required"`
	Password  string         `json:"password" validate:"required"`
}

func FromDomain(domain profiles.Domain) Profile {
	return Profile{
		ID:        domain.ID,
		Name:      domain.Name,
		Address:   domain.Address,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
