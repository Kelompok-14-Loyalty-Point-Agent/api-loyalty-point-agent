package response

import (
	redeems "api-loyalty-point-agent/businesses/redeems"
	"time"
)

type Redeem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	Phone     string    `json:"product"`
	Cost      uint      `json:"cost"`
}

func FromDomain(domain redeems.Domain) Redeem {
	return Redeem{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		Phone:     domain.Phone,
		Cost:      domain.Cost,
	}
}
