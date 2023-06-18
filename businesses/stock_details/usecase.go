package stock_details

import (
	"api-loyalty-point-agent/app/middlewares"
	"context"
)

type stock_detailUsecase struct {
	stock_detailRepository Repository
	jwtAuth         *middlewares.JWTConfig
}

func NewStockDetailUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &stock_detailUsecase{
		stock_detailRepository: repository,
		jwtAuth:         jwtAuth,
	}
}

func (usecase *stock_detailUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.stock_detailRepository.GetAll(ctx)
}

func (usecase *stock_detailUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.stock_detailRepository.GetByID(ctx, id)
}

func (usecase *stock_detailUsecase) Create(ctx context.Context, stock_detailDomain *Domain) (Domain, error) {
	return usecase.stock_detailRepository.Create(ctx, stock_detailDomain)
}

func (usecase *stock_detailUsecase) Update(ctx context.Context, stock_detailDomain *Domain, id string) (Domain, error) {
	return usecase.stock_detailRepository.Update(ctx, stock_detailDomain, id)
}

func (usecase *stock_detailUsecase) Delete(ctx context.Context, id string) error {
	return usecase.stock_detailRepository.Delete(ctx, id)
}

func (usecase *stock_detailUsecase) GetAllByStockID(ctx context.Context, stockid string) ([]Domain, error) {
	return usecase.stock_detailRepository.GetAllByStockID(ctx, stockid)
}
