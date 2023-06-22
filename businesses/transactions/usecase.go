package transactions

import (
	"api-loyalty-point-agent/app/middlewares"
	"context"
)

type transactionUsecase struct {
	transactionRepository Repository
	jwtAuth               *middlewares.JWTConfig
}

func NewTransactionUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &transactionUsecase{
		transactionRepository: repository,
		jwtAuth:               jwtAuth,
	}
}

func (usecase *transactionUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.transactionRepository.GetAll(ctx)
}

func (usecase *transactionUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.transactionRepository.GetByID(ctx, id)
}

func (usecase *transactionUsecase) Create(ctx context.Context, transactionDomain *Domain) (Domain, error) {
	return usecase.transactionRepository.Create(ctx, transactionDomain)
}

func (usecase *transactionUsecase) GetAllByUserID(ctx context.Context, userid string) ([]Domain, error) {
	return usecase.transactionRepository.GetAllByUserID(ctx, userid)
}

func (usecase *transactionUsecase) UpdatePoint(ctx context.Context, transactionDomain *Domain, id string) (Domain, error) {
	return usecase.transactionRepository.UpdatePoint(ctx, transactionDomain, id)
}

// func (usecase *transactionUsecase) GetTotalTransactionMade(ctx context.Context, userid string) (Domain, error) {
// 	return usecase.transactionRepository.GetTotalTransactionMade(ctx, userid)
// }

// func (usecase *transactionUsecase) GetTotalTransactionMade(ctx context.Context) (map[string]int, error) {
// 	return usecase.transactionRepository.GetTotalTransactionMade(ctx, totalTransactionMap, id)
// }
