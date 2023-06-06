package stocks

import (
	"api-loyalty-point-agent/businesses/stock_transactions"
	"api-loyalty-point-agent/app/middlewares"
	"context"
)

type stockUsecase struct {
	stockRepository Repository
	jwtAuth         *middlewares.JWTConfig
}

func NewStockUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &stockUsecase{
		stockRepository: repository,
		jwtAuth:         jwtAuth,
	}
}

func (usecase *stockUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.stockRepository.GetAll(ctx)
}

func (usecase *stockUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.stockRepository.GetByID(ctx, id)
}

func (usecase *stockUsecase) Create(ctx context.Context, stockDomain *Domain) (Domain, error) {
	return usecase.stockRepository.Create(ctx, stockDomain)
}

func (usecase *stockUsecase) Update(ctx context.Context, stockDomain *Domain, id string) (Domain, error) {
	return usecase.stockRepository.Update(ctx, stockDomain, id)
}

func (usecase *stockUsecase) Delete(ctx context.Context, id string) error {
	return usecase.stockRepository.Delete(ctx, id)
}

func (usecase *stockUsecase) AddStock(ctx context.Context, stock_transactionDomain *stock_transactions.Domain) (stock_transactions.Domain, error) {
	return usecase.stockRepository.AddStock(ctx, stock_transactionDomain)
}
