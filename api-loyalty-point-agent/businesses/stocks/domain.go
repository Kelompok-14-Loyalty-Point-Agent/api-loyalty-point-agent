package stocks

import (
	"api-loyalty-point-agent/businesses/providers"
	"api-loyalty-point-agent/businesses/stock_transactions"
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	Type       string
	TotalStock float64
	LastTopUp  time.Time
	Provider   providers.Domain
	ProviderID uint
}
type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	AddStock(ctx context.Context, stock_transactionDomain *stock_transactions.Domain) (stock_transactions.Domain, error)
}
type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	AddStock(ctx context.Context, stock_transactionDomain *stock_transactions.Domain) (stock_transactions.Domain, error)
}
