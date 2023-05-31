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

func (usecase *providerUsecase) GetAllProvider(ctx context.Context) ([]Domain, error) {
	return usecase.providerRepository.GetAllProvider(ctx)
}

func (usecase *providerUsecase) GetByIDProvider(ctx context.Context, id string) (Domain, error) {
	return usecase.providerRepository.GetByIDProvider(ctx, id)
}

func (usecase *providerUsecase) CreateProvider(ctx context.Context, providerDomain *Domain) (Domain, error) {
	return usecase.providerRepository.CreateProvider(ctx, providerDomain)
}

func (usecase *providerUsecase) UpdateProvider(ctx context.Context, providerDomain *Domain, id string) (Domain, error) {
	return usecase.providerRepository.UpdateProvider(ctx, providerDomain, id)
}

func (usecase *providerUsecase) DeleteProvider(ctx context.Context, id string) error {
	return usecase.providerRepository.DeleteProvider(ctx, id)
}
