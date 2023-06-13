package providers

import (
	"api-loyalty-point-agent/app/middlewares"
	"context"
)

type providerUsecase struct {
	providerRepository Repository
	jwtAuth            *middlewares.JWTConfig
}

func NewProviderUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &providerUsecase{
		providerRepository: repository,
		jwtAuth:            jwtAuth,
	}
}

func (usecase *providerUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.providerRepository.GetAll(ctx)
}

func (usecase *providerUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.providerRepository.GetByID(ctx, id)
}
