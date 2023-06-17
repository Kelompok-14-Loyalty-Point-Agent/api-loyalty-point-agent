package profiles

import (
	"api-loyalty-point-agent/app/middlewares"
	"context"
)

type profileUsecase struct {
	profileRepository Repository
	jwtAuth           *middlewares.JWTConfig
}

func NewProfileUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &profileUsecase{
		profileRepository: repository,
		jwtAuth:           jwtAuth,
	}
}

func (usecase *profileUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.profileRepository.GetAll(ctx)
}

func (usecase *profileUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.profileRepository.GetByID(ctx, id)
}

func (usecase *profileUsecase) Update(ctx context.Context, profileDomain *Domain, id string) (Domain, error) {
	// Periksa apakah profil sedang dalam mode edit atau tidak
	existingProfile, err := usecase.profileRepository.GetByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	if existingProfile.IsEditing {
		// Jika profil sedang dalam mode edit, lakukan update data
		profileDomain, err := usecase.profileRepository.Update(ctx, profileDomain, id)
		if err != nil {
			return Domain{}, err
		}
		return profileDomain, nil
	} else {
		// Jika profil tidak dalam mode edit, hanya update password
		profileDomain := &Domain{
			ID:       existingProfile.ID,
			Name:     existingProfile.Name,
			Address:  existingProfile.Address,
			Password: profileDomain.Password,
		}

		_, err := usecase.profileRepository.Update(ctx, profileDomain, id)
		if err != nil {
			return Domain{}, err
		}
		return *profileDomain, nil
	}
}

func (usecase *profileUsecase) Delete(ctx context.Context, id string) error {
	return usecase.profileRepository.Delete(ctx, id)
}
