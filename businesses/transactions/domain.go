package transactions

import (
	"api-loyalty-point-agent/businesses/stock_details"
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
	Product        string
	Payment_method string
	Point          float64
	Status         string
	Description    string
	UserID         uint
	UserName       string
	Member         string
	URL            string
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, transactionDomain *Domain) (Domain, error)
	GetAllByUserID(ctx context.Context, id string) ([]Domain, error)
	GetAllByUserIDSorted(ctx context.Context, id string) ([]Domain, error)
	UpdatePoint(ctx context.Context, transactionDomain *Domain, id string) (Domain, error)
}
type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, transactionDomain *Domain) (Domain, error)
	GetAllByUserID(ctx context.Context, userid string) ([]Domain, error)
	GetAllByUserIDSorted(ctx context.Context, id string) ([]Domain, error)
	UpdatePoint(ctx context.Context, transactionDomain *Domain, id string) (Domain, error)
}
