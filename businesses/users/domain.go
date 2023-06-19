package users

import (
	"api-loyalty-point-agent/businesses/profiles"
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
	Role      string
	ProfileID uint
	Profile   profiles.Domain
}
type Usecase interface {
	GetAllCustomers(ctx context.Context) ([]Domain, error)
	Register(ctx context.Context, userDomain *Domain) (Domain, error)
	Login(ctx context.Context, userDomain *Domain) (string, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	UpdateProfileCustomer(ctx context.Context, userDomain *Domain, id string) (Domain, error)
	UpdateProfileAdmin(ctx context.Context, userDomain *Domain, id string) (Domain, error)
	ChangePassword(ctx context.Context, userDomain *Domain, id string) (Domain, error)
	DeleteCustomer(ctx context.Context, id string) (error)
}
type Repository interface {
	GetAllCustomers(ctx context.Context) ([]Domain, error)
	Register(ctx context.Context, userDomain *Domain) (Domain, error)
	GetByEmail(ctx context.Context, userDomain *Domain) (Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	UpdateProfileCustomer(ctx context.Context, userDomain *Domain, id string) (Domain, error)
	UpdateProfileAdmin(ctx context.Context, userDomain *Domain, id string) (Domain, error)
	ChangePassword(ctx context.Context, userDomain *Domain, id string) (Domain, error)
	DeleteCustomer(ctx context.Context, id string) (error)
}
