package profiles

import (
	"api-loyalty-point-agent/businesses/profiles"
	"context"

	"gorm.io/gorm"
)

type profileRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) profiles.Repository {
	return &profileRepository{
		conn: conn,
	}
}

func (pr *profileRepository) GetAll(ctx context.Context) ([]profiles.Domain, error) {
	var records []Profile

	if err := pr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	profiles := []profiles.Domain{}

	for _, profile := range records {
		profiles = append(profiles, profile.ToDomain())
	}

	return profiles, nil
}

func (pr *profileRepository) GetByID(ctx context.Context, id string) (profiles.Domain, error) {
	var profile Profile

	if err := pr.conn.WithContext(ctx).First(&profile, "id = ?", id).Error; err != nil {
		return profiles.Domain{}, err
	}

	return profile.ToDomain(), nil
}

func (pr *profileRepository) Update(ctx context.Context, profileDomain *profiles.Domain, id string) (profiles.Domain, error) {
	profile, err := pr.GetByID(ctx, id)
	if err != nil {
		return profiles.Domain{}, err
	}

	updatedProfile := FromDomain(&profile)
	// Update data sesuai dengan profileDomain
	updatedProfile.Address = profileDomain.Address
	// updatedProfile.Password = profileDomain.Password
	// updatedProfile.TransactionMade = profileDomain.TransactionMade + 1

	if err := pr.conn.WithContext(ctx).Save(&updatedProfile).Error; err != nil {
		return profiles.Domain{}, err
	}

	return updatedProfile.ToDomain(), nil
}

func (pr *profileRepository) Delete(ctx context.Context, id string) error {
	profile, err := pr.GetByID(ctx, id)
	if err != nil {
		return err
	}

	deletedprofile := FromDomain(&profile)

	if err := pr.conn.WithContext(ctx).Unscoped().Delete(&deletedprofile).Error; err != nil {
		return err
	}

	return nil
}
