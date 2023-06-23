package redeems

import (
	// "api-loyalty-point-agent/businesses/users"

	"context"
	"time"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	Phone     string
	Cost      uint
	//bukan
	UserID uint
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
