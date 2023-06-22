package users

import (
	"api-loyalty-point-agent/app/middlewares"
	"context"
)

type userUsecase struct {
	userRepository Repository
	jwtAuth            *middlewares.JWTConfig
}

func NewUserUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &userUsecase{
		userRepository: repository,
		jwtAuth:            jwtAuth,
	}
}

func (usecase *userUsecase) GetAllCustomers(ctx context.Context) ([]Domain, error) {
	return usecase.userRepository.GetAllCustomers(ctx)
}

func (usecase *userUsecase) Register(ctx context.Context, userDomain *Domain) (Domain, error) {
	return usecase.userRepository.Register(ctx, userDomain)
}

func (usecase *userUsecase) Login(ctx context.Context, userDomain *Domain) (string, error) {
	user, err := usecase.userRepository.GetByEmail(ctx, userDomain)

	if err != nil {
		return "", err
	}

	token, err := usecase.jwtAuth.GenerateToken(int(user.ID), user.Role)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (usecase *userUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.userRepository.GetByID(ctx, id)
}

func (usecase *userUsecase) UpdateProfileCustomer(ctx context.Context, userDomain *Domain, id string) (Domain, error) {
	return usecase.userRepository.UpdateProfileCustomer(ctx, userDomain, id)
}

func (usecase *userUsecase) UpdateProfileAdmin(ctx context.Context, userDomain *Domain, id string) (Domain, error) {
	return usecase.userRepository.UpdateProfileAdmin(ctx, userDomain, id)
}

func (usecase *userUsecase) ChangePassword(ctx context.Context, userDomain *Domain, id string) (Domain, error) {
	return usecase.userRepository.ChangePassword(ctx, userDomain, id)
}

func (usecase *userUsecase) ChangePicture(ctx context.Context, filename string, id string) (string, string, error) {
	return usecase.userRepository.ChangePicture(ctx, filename, id)
}

func (usecase *userUsecase) DeleteCustomer(ctx context.Context, id string) (error) {
	return usecase.userRepository.DeleteCustomer(ctx, id)
}