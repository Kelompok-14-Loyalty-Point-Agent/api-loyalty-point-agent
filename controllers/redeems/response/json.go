package response

import (
	"api-loyalty-point-agent/businesses/redeems"
	"time"

	"gorm.io/gorm"
)

type Redeem struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	Product       string         `json:"product"`
	PaymentMethod string         `json:"payment_method"`
	DateExchange  time.Time      `json:"date_exchange"`
	Cost          float64        `json:"cost"`
	UserID        uint           `json:"user_id"`
}

func FromDomain(domain redeems.Domain) Redeem {
	return Redeem{
		ID:            domain.ID,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
		DeletedAt:     domain.DeletedAt,
		Product:       domain.Product,
		PaymentMethod: domain.Payment_method,
		DateExchange:  domain.DateExchange,
		Cost:          domain.Cost,
		UserID:        domain.UserID,
	}
}
