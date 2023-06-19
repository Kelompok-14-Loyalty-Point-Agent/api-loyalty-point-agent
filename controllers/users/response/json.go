package response

import (
	"api-loyalty-point-agent/businesses/profiles"
	"api-loyalty-point-agent/businesses/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint            `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt gorm.DeletedAt  `json:"deleted_at"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Password  string          `json:"password"`
	Role      string          `json:"role"`
	ProfileID uint            `json:"profile_id"`
	Profile   profiles.Domain `json:"profile"`
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		Role:      domain.Role,
		ProfileID: domain.ProfileID,
		Profile:   domain.Profile,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
