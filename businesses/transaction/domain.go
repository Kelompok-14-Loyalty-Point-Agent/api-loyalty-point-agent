package transaction

import (
	"api-loyalty-point-agent/businesses/stocks"
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	Phone          string
	StockID        uint
	Stock          stocks.Domain
	Product        uint
	Payment_method uint
	Point          uint
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, transactionDomain *Domain) (Domain, error)
}
type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, transactionDomain *Domain) (Domain, error)
}
