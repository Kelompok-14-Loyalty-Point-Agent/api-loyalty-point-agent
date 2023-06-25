package redeems

import (
	"api-loyalty-point-agent/businesses/redeems"

	"time"

	"gorm.io/gorm"
)

type Redeem struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	DateExchange  time.Time      `json:"date_exhange"`
	PaymentMethod string         `json:"payment_method"`
	Product       string         `json:"product"`
	Phone         string         `json:"phone"`
	Cost          float32        `json:"cost"`
	UserID        uint           `json:"user_id"`
	VoucherID     uint           `json:"voucher_id"`
}

func (rec *Redeem) ToDomain() redeems.Domain {
	return redeems.Domain{
		ID:             rec.ID,
		CreatedAt:      rec.CreatedAt,
		UpdatedAt:      rec.UpdatedAt,
		DeletedAt:      rec.DeletedAt,
		DateExchange:   rec.DateExchange,
		Payment_method: rec.PaymentMethod,
		Product:        rec.Product,
		Phone:          rec.Phone,
		Cost:           rec.Cost,
		UserID:         rec.UserID,
		VoucherID:      rec.VoucherID,
	}
}

func FromDomain(domain *redeems.Domain) *Redeem {
	return &Redeem{
		ID:            domain.ID,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
		DeletedAt:     domain.DeletedAt,
		DateExchange:  domain.DateExchange,
		PaymentMethod: domain.Payment_method,
		Product:       domain.Product,
		Phone:         domain.Phone,
		Cost:          domain.Cost,
		UserID:        domain.UserID,
		VoucherID:     domain.VoucherID,
	}
}
