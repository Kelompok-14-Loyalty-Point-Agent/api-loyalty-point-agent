package transactions

import (
	stock_details "api-loyalty-point-agent/businesses/stock_details"

	// stock_details "api-loyalty-point-agent/businesses/Stock_details"
	// "api-loyalty-point-agent/businesses/stocks"

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
	StockDetailsID uint
	StockDetails   stock_details.Domain
	StockID        uint
	Price          float64
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
