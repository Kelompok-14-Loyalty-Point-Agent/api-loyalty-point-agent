package vouchers

import (
	"context"
)

type Domain struct {
	ID      uint
	Product string
	Benefit string
	Cost    uint
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	// GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, voucherDomain *Domain) (Domain, error)
}
type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	// GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, voucherDomain *Domain) (Domain, error)
}
