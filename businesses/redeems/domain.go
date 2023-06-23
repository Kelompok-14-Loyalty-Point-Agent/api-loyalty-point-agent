package redeems

import (
	// "api-loyalty-point-agent/businesses/users"

	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	Product        string
	Payment_method string
	Phone          string
	DateExchange   time.Time
	Cost           float32
	UserID         uint
	VoucherID      uint
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)

}
type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
}
