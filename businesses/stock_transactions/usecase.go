package stock_transactions

import (
	"api-loyalty-point-agent/app/middlewares"
	"context"
)

type stock_transactionUsecase struct {
	stock_transactionRepository Repository
	jwtAuth         *middlewares.JWTConfig
}

func NewStockTransactionUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &stock_transactionUsecase{
		stock_transactionRepository: repository,
		jwtAuth:         jwtAuth,
	}
}

func (usecase *stock_transactionUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.stock_transactionRepository.GetAll(ctx)
}

func (usecase *stock_transactionUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.stock_transactionRepository.GetByID(ctx, id)
}
