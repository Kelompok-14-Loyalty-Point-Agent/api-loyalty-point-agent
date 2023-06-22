package voucher

import (
	"context"
)

type Domain struct {
	ID    uint
	Title string
	Url   string
	Cost  uint
	Point uint
	// Expired	uint
	// Date_Exhange time.Time
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
