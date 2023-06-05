package user_details

import (
	"api-loyalty-point-agent/businesses/user_details"
	"context"

	"gorm.io/gorm"
)

type user_detailRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) user_details.Repository {
	return &user_detailRepository{
		conn: conn,
	}
}

func (cr *user_detailRepository) GetAll(ctx context.Context) ([]user_details.Domain, error) {
	var records []UserDetail

	if err := cr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	user_details := []user_details.Domain{}

	for _, user_detail := range records {
		user_details = append(user_details, user_detail.ToDomain())
	}

	return user_details, nil
}

func (cr *user_detailRepository) GetByID(ctx context.Context, id string) (user_details.Domain, error) {
	var user_detail UserDetail

	if err := cr.conn.WithContext(ctx).First(&user_detail, "id = ?", id).Error; err != nil {
		return user_details.Domain{}, err
	}

	return user_detail.ToDomain(), nil
}

func (cr *user_detailRepository) Update(ctx context.Context, user_detailDomain *user_details.Domain, id string) (user_details.Domain, error) {
	user_detail, err := cr.GetByID(ctx, id)

	if err != nil {
		return user_details.Domain{}, err
	}

	updatedUserDetail := FromDomain(&user_detail)

	if updatedUserDetail.Age != user_detailDomain.Age {
		updatedUserDetail.Age = user_detailDomain.Age
	}

	if updatedUserDetail.Gender != user_detailDomain.Gender {
		updatedUserDetail.Gender = user_detailDomain.Gender
	}
	
	if updatedUserDetail.Address != user_detailDomain.Address {
		updatedUserDetail.Address = user_detailDomain.Address
	}

	if updatedUserDetail.PhoneNumber != user_detailDomain.PhoneNumber {
		updatedUserDetail.PhoneNumber = user_detailDomain.PhoneNumber
	}

	if err := cr.conn.WithContext(ctx).Save(&updatedUserDetail).Error; err != nil {
		return user_details.Domain{}, err
	}

	return updatedUserDetail.ToDomain(), nil
}

func (cr *user_detailRepository) Delete(ctx context.Context, id string) error {
	user_detail, err := cr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedUserDetail := FromDomain(&user_detail)

	if err := cr.conn.WithContext(ctx).Unscoped().Delete(&deletedUserDetail).Error; err != nil {
		return err
	}

	return nil
}

