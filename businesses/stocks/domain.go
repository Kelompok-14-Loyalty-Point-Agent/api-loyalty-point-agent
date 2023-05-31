package stocks

import (
	"api-loyalty-point-agent/businesses/providers"
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	Name       string
	Type       string
	Stock      float64
	Price      float64
	Quantity   float64
	Provider   providers.Domain
	ProviderID uint
}
type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, stockDomain *Domain) (Domain, error)
	Update(ctx context.Context, stockDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
}
type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, stockDomain *Domain) (Domain, error)
	Update(ctx context.Context, stockDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
}
