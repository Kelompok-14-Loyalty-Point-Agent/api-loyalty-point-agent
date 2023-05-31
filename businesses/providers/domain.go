package providers

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
	Name      string
	URL       string
}
type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, providerDomain *Domain) (Domain, error)
	Update(ctx context.Context, providerDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, providerDomain *Domain) (Domain, error)
	Update(ctx context.Context, providerDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
}
