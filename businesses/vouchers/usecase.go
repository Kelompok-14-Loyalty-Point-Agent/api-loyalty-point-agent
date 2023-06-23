package vouchers

import (
	"api-loyalty-point-agent/app/middlewares"
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

// func (usecase *voucherUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
// 	return usecase.voucherRepository.GetByID(ctx, id)
// }

func (usecase *voucherUsecase) Create(ctx context.Context, transactionDomain *Domain) (Domain, error) {
	return usecase.voucherRepository.Create(ctx, transactionDomain)
}
