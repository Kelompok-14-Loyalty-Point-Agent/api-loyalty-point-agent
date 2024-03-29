package stock_details

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Stock     float64
	Price     float64
	Quantity  float64
	StockID   uint
}
type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, stock_detailDomain *Domain) (Domain, error)
	Update(ctx context.Context, stock_detailDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
	GetAllByStockID(ctx context.Context, stockid string) ([]Domain, error)
}
type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, stock_detailDomain *Domain) (Domain, error)
	Update(ctx context.Context, stock_detailDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
	GetAllByStockID(ctx context.Context, stockid string) ([]Domain, error)
}
