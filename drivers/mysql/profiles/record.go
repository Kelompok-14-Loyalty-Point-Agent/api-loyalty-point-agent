package profiles

import (
	"api-loyalty-point-agent/businesses/profiles"
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	Password  string         `json:"password"`
}

func (record *Profile) ToDomain() profiles.Domain {
	return profiles.Domain{
		ID:        record.ID,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		DeletedAt: record.DeletedAt,
		Name:      record.Name,
		Address:   record.Address,
		Password:  record.Password,
	}
}

func FromDomain(domain *profiles.Domain) *Profile {
	return &Profile{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Name:      domain.Name,
		Address:   domain.Address,
		Password:  domain.Password,
	}
}
