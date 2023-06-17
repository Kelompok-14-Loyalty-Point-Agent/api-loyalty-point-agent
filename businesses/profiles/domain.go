package profiles

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
	Address   string
	Password  string
	IsEditing bool
}
type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	Update(ctx context.Context, profileDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
}
type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	Update(ctx context.Context, profileDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
}
