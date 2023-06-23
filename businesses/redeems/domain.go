package redeems

import (
	"context"
	"time"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	Phone     string
	Cost      uint
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	// GetByID(ctx context.Context, id string) (Domain, error)
	RedeemVoucher(ctx context.Context, voucherDomain *Domain) (Domain, error)
}
type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	// GetByID(ctx context.Context, id string) (Domain, error)
	RedeemVoucher(ctx context.Context, voucherDomain *Domain) (Domain, error)
}
