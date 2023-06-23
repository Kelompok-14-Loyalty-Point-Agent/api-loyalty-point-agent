package vouchers

import (
	"api-loyalty-point-agent/businesses/redeems"
	"context"
)

type Domain struct {
	ID      uint
	Product string
	Benefit string
	Cost    float32
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	RedeemVoucher(ctx context.Context, redeemDomain *redeems.Domain) (redeems.Domain, error)
}
type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	RedeemVoucher(ctx context.Context, redeemDomain *redeems.Domain) (redeems.Domain, error)
}
