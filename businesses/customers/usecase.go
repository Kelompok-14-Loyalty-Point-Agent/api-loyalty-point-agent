package customers

import (
	"api-loyalty-point-agent/app/middlewares"
	"context"
)

type customerUsecase struct {
	customerRepository Repository
	jwtAuth            *middlewares.JWTConfig
}

func NewCustomerUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &customerUsecase{
		customerRepository: repository,
		jwtAuth:            jwtAuth,
	}
}

func (usecase *customerUsecase) GetAllCustomers(ctx context.Context) ([]Domain, error) {
	return usecase.customerRepository.GetAllCustomers(ctx)
}

func (usecase *customerUsecase) Register(ctx context.Context, customerDomain *Domain) (Domain, error) {
	return usecase.customerRepository.Register(ctx, customerDomain)
}

func (usecase *customerUsecase) Login(ctx context.Context, customerDomain *Domain) (string, error) {
	customer, err := usecase.customerRepository.GetByEmail(ctx, customerDomain)

	if err != nil {
		return "", err
	}

	token, err := usecase.jwtAuth.GenerateToken(int(customer.ID))

	if err != nil {
		return "", err
	}

	return token, nil
}
