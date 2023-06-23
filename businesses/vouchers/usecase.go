package vouchers

import (
	"api-loyalty-point-agent/app/middlewares"
	"api-loyalty-point-agent/businesses/redeems"
	"context"
)

type voucherUsecase struct {
	voucherRepository Repository
	jwtAuth           *middlewares.JWTConfig
}

func NewVoucherUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &voucherUsecase{
		voucherRepository: repository,
		jwtAuth:           jwtAuth,
	}
}

func (usecase *voucherUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.voucherRepository.GetAll(ctx)
}

func (usecase *voucherUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.voucherRepository.GetByID(ctx, id)
}

func (usecase *voucherUsecase) RedeemVoucher(ctx context.Context, redeemDomain *redeems.Domain) (redeems.Domain, error) {
	return usecase.voucherRepository.RedeemVoucher(ctx, redeemDomain)
}
