package redeems

import (
	"api-loyalty-point-agent/businesses/redeems"
	"time"
)

type Redeem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"create_at"`
	Phone     string    `json:"phone"`
	Cost      uint      `json:"cost"`
}

func (rec *Redeem) ToDomain() redeems.Domain {
	return redeems.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		Phone:     rec.Phone,
		Cost:      rec.Cost,
	}
}

func FromDomain(domain *redeems.Domain) *Redeem {
	return &Redeem{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		Phone:     domain.Phone,
		Cost:      domain.Cost,
	}
}
