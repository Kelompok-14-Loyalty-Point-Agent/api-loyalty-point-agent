package redeems

import (
	"api-loyalty-point-agent/businesses/redeems"

	"context"

	"gorm.io/gorm"
)

type redeemRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) redeems.Repository {
	return &redeemRepository{
		conn: conn,
	}
}

func (cr *redeemRepository) GetAll(ctx context.Context) ([]redeems.Domain, error) {
	var records []Redeem

	if err := cr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	redeems := []redeems.Domain{}

	for _, redeem := range records {
		redeems = append(redeems, redeem.ToDomain())
	}

	return redeems, nil
}

func (ur *redeemRepository) GetByID(ctx context.Context, id string) (redeems.Domain, error) {
	var redeem Redeem

	if err := ur.conn.WithContext(ctx).First(&redeem, "id = ?", id).Error; err != nil {
		return redeems.Domain{}, err
	}

	return redeem.ToDomain(), nil

}




