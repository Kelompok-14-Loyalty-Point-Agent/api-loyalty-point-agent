package stock_transactions

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID            uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	ProviderName  string
	InputStock    float64
	PayAmount     float64
	PaymentMethod string
	Status        string
	StockID       uint
	UserID        uint
}
type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
}
type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
}
