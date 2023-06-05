package user_details

import (
	"api-loyalty-point-agent/app/middlewares"
	"context"
)

type user_detailUsecase struct {
	user_detailRepository Repository
	jwtAuth         *middlewares.JWTConfig
}

func NewUserDetailUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &user_detailUsecase{
		user_detailRepository: repository,
		jwtAuth:         jwtAuth,
	}
}

func (usecase *user_detailUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.user_detailRepository.GetAll(ctx)
}

func (usecase *user_detailUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.user_detailRepository.GetByID(ctx, id)
}

func (usecase *user_detailUsecase) Update(ctx context.Context, user_detailDomain *Domain, id string) (Domain, error) {
	return usecase.user_detailRepository.Update(ctx, user_detailDomain, id)
}

func (usecase *user_detailUsecase) Delete(ctx context.Context, id string) error {
	return usecase.user_detailRepository.Delete(ctx, id)
}