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
	GetAllProvider(ctx context.Context) ([]Domain, error)
	GetByIDProvider(ctx context.Context, id string) (Domain, error)
	CreateProvider(ctx context.Context, providerDomain *Domain) (Domain, error)
	UpdateProvider(ctx context.Context, providerDomain *Domain, id string) (Domain, error)
	DeleteProvider(ctx context.Context, id string) error
}

type Repository interface {
	GetAllProvider(ctx context.Context) ([]Domain, error)
	GetByIDProvider(ctx context.Context, id string) (Domain, error)
	CreateProvider(ctx context.Context, providerDomain *Domain) (Domain, error)
	UpdateProvider(ctx context.Context, providerDomain *Domain, id string) (Domain, error)
	DeleteProvider(ctx context.Context, id string) error
}
