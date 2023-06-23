package redeems

import (
	"api-loyalty-point-agent/app/middlewares"
	"context"
)

type redeemUsecase struct {
	redeemRepository Repository
	jwtAuth          *middlewares.JWTConfig
}

func NewRedeemUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &redeemUsecase{
		redeemRepository: repository,
		jwtAuth:          jwtAuth,
	}
}

func (usecase *redeemUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.redeemRepository.GetAll(ctx)
}

// func (usecase *redeemUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
// 	return usecase.redeemRepository.GetByID(ctx, id)
// }

func (usecase *redeemUsecase) RedeemVoucher(ctx context.Context, transactionDomain *Domain) (Domain, error) {
	return usecase.redeemRepository.RedeemVoucher(ctx, transactionDomain)
}
