package customers

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
	Email     string
	Password  string
}
type Usecase interface {
	GetAllCustomers(ctx context.Context) ([]Domain, error)
	Register(ctx context.Context, userDomain *Domain) (Domain, error)
	Login(ctx context.Context, userDomain *Domain) (string, error)
}
type Repository interface {
	GetAllCustomers(ctx context.Context) ([]Domain, error)
	Register(ctx context.Context, userDomain *Domain) (Domain, error)
	GetByEmail(ctx context.Context, userDomain *Domain) (Domain, error)
}
