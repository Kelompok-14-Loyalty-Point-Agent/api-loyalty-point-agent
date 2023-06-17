package providers

import (
	"context"
	"api-loyalty-point-agent/businesses/providers"

	"gorm.io/gorm"
)

type providerRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) providers.Repository {
	return &providerRepository{
		conn: conn,
	}
}

func (cr *providerRepository) GetAll(ctx context.Context) ([]providers.Domain, error) {
	var records []Provider

	if err := cr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	providers := []providers.Domain{}

	for _, provider := range records {
		providers = append(providers, provider.ToDomain())
	}

	return providers, nil
}

func (cr *providerRepository) GetByID(ctx context.Context, id string) (providers.Domain, error) {
	var provider Provider

	if err := cr.conn.WithContext(ctx).First(&provider, "id = ?", id).Error; err != nil {
		return providers.Domain{}, err
	}

	return provider.ToDomain(), nil
}

func (cr *providerRepository) Create(ctx context.Context, providerDomain *providers.Domain) (providers.Domain, error) {
	record := FromDomain(providerDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return providers.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return providers.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (cr *providerRepository) Update(ctx context.Context, providerDomain *providers.Domain, id string) (providers.Domain, error) {
	provider, err := cr.GetByID(ctx, id)

	if err != nil {
		return providers.Domain{}, err
	}

	updatedProvider := FromDomain(&provider)

	updatedProvider.Name = providerDomain.Name

	if err := cr.conn.WithContext(ctx).Save(&updatedProvider).Error; err != nil {
		return providers.Domain{}, err
	}

	return updatedProvider.ToDomain(), nil
}

func (cr *providerRepository) Delete(ctx context.Context, id string) error {
	provider, err := cr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedProvider := FromDomain(&provider)

	if err := cr.conn.WithContext(ctx).Unscoped().Delete(&deletedProvider).Error; err != nil {
		return err
	}

	return nil
}