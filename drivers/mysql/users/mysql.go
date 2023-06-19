package users

import (
	"api-loyalty-point-agent/businesses/users"
	"api-loyalty-point-agent/drivers/mysql/profiles"

	"context"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) GetAllCustomers(ctx context.Context) ([]users.Domain, error) {
	var records []User

	if err := ur.conn.WithContext(ctx).Preload("Profile").Find(&records, "role = ?", "customer").Error; err != nil {
		return nil, err
	}

	users := []users.Domain{}

	for _, user := range records {
		users = append(users, user.ToDomain())
	}

	return users, nil
}

func (ur *userRepository) Register(ctx context.Context, userDomain *users.Domain) (users.Domain, error) {
	var profile profiles.Profile

	result := ur.conn.WithContext(ctx).Create(&profile)

	if err := result.Error; err != nil {
		return users.Domain{}, err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)

	if err != nil {
		return users.Domain{}, err
	}

	record := FromDomain(userDomain)

	record.Password = string(password)
	record.Profile = profile
	record.ProfileID = profile.ID

	result = ur.conn.WithContext(ctx).Preload("Profile").Create(&record)

	if err := result.Error; err != nil {
		return users.Domain{}, err
	}

	err = result.Last(&record).Error

	if err != nil {
		return users.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (ur *userRepository) GetByEmail(ctx context.Context, userDomain *users.Domain) (users.Domain, error) {
	var user User

	err := ur.conn.WithContext(ctx).First(&user, "email = ?", userDomain.Email).Error

	if err != nil {
		return users.Domain{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDomain.Password))

	if err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}

func (ur *userRepository) GetByID(ctx context.Context, id string) (users.Domain, error) {
	var user User

	if err := ur.conn.WithContext(ctx).Preload("Profile").First(&user, "id = ?", id).Error; err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil

}

func (ur *userRepository) UpdateProfileCustomer(ctx context.Context, userDomain *users.Domain, id string) (users.Domain, error) {
	var user User

	if err := ur.conn.WithContext(ctx).Preload("Profile").First(&user, "id = ?", id).Error; err != nil {
		return users.Domain{}, err
	}

	if user.Name != userDomain.Name {
		user.Name = userDomain.Name
	}

	if user.Email != user.Email {
		user.Email = userDomain.Email
	}

	var profile profiles.Profile

	if err := ur.conn.WithContext(ctx).First(&profile, "id = ?", id).Error; err != nil {
		return users.Domain{}, err
	}

	updatedProfiles := profile.ToDomain()

	updatedProfiles.Address = profile.Address
	updatedProfiles.Age = profile.Age
	updatedProfiles.Gender = profile.Gender
	updatedProfiles.Point = profile.Point
	updatedProfiles.Phone = profile.Phone
	updatedProfiles.TransactionMade = profile.TransactionMade
	updatedProfiles.TotalRedeem = profile.TotalRedeem
	updatedProfiles.MonthlyTransaction = profile.MonthlyTransaction
	updatedProfiles.Member = profile.Member

	if user.Email != userDomain.Email {
		user.Email = userDomain.Email
	}

	if profile.Address != userDomain.Profile.Address {
		profile.Address = userDomain.Profile.Address
		updatedProfiles.Address = profile.Address
	}

	if profile.Phone != userDomain.Profile.Phone {
		profile.Phone = userDomain.Profile.Phone
		updatedProfiles.Phone = profile.Phone
	}

	if profile.Age != userDomain.Profile.Age {
		profile.Age = userDomain.Profile.Age
		updatedProfiles.Age = profile.Age
	}

	if profile.Gender != userDomain.Profile.Gender {
		profile.Gender = userDomain.Profile.Gender
		updatedProfiles.Gender = profile.Gender
	}

	if err := ur.conn.WithContext(ctx).Save(&profile).Error; err != nil {
		return users.Domain{}, err
	}

	if err := ur.conn.WithContext(ctx).Save(&user).Error; err != nil {
		return users.Domain{}, err
	}

	user.Profile = *profiles.FromDomain(&updatedProfiles)

	return user.ToDomain(), nil
}

func (ur *userRepository) UpdateProfileAdmin(ctx context.Context, userDomain *users.Domain, id string) (users.Domain, error) {
	var user User

	if err := ur.conn.WithContext(ctx).Preload("Profile").First(&user, "id = ?", id).Error; err != nil {
		return users.Domain{}, err
	}

	if user.Name != userDomain.Name {
		user.Name = userDomain.Name
	}

	var profile profiles.Profile

	if err := ur.conn.WithContext(ctx).First(&profile, "id = ?", id).Error; err != nil {
		return users.Domain{}, err
	}

	updatedProfiles := profile.ToDomain()

	updatedProfiles.Address = profile.Address

	if user.Name != userDomain.Name {
		user.Name = userDomain.Name
	}

	if profile.Address != userDomain.Profile.Address {
		profile.Address = userDomain.Profile.Address
		updatedProfiles.Address = profile.Address
	}

	if err := ur.conn.WithContext(ctx).Save(&profile).Error; err != nil {
		return users.Domain{}, err
	}

	if err := ur.conn.WithContext(ctx).Save(&user).Error; err != nil {
		return users.Domain{}, err
	}

	user.Profile = *profiles.FromDomain(&updatedProfiles)

	return user.ToDomain(), nil
}


func (ur *userRepository) ChangePassword(ctx context.Context, userDomain *users.Domain, id string) (users.Domain, error) {
	user, err := ur.GetByID(ctx, id)
	if err != nil {
		return users.Domain{}, err
	}

	updatedUser := FromDomain(&user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDomain.Password))
	if err := ur.conn.WithContext(ctx).Save(&updatedUser).Error; err != nil {
		return users.Domain{}, err
	}
	
	newPassword, err := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)


	if err != nil {
		updatedUser.Password = string(newPassword)
	}

	if err := ur.conn.WithContext(ctx).Save(&updatedUser).Error; err != nil {
		return users.Domain{}, err
	}

	return updatedUser.ToDomain(), nil
}

func (ur *userRepository) DeleteCustomer(ctx context.Context, id string) error {
	user, err := ur.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedUser := FromDomain(&user)

	if err := ur.conn.WithContext(ctx).Unscoped().Delete(&deletedUser).Error; err != nil {
		return err
	}

	var deletedProfile profiles.Profile

	if err := ur.conn.WithContext(ctx).First(&deletedProfile, "id = ?", id).Error; err != nil {
		return err
	}

	if err := ur.conn.WithContext(ctx).Unscoped().Delete(&deletedProfile).Error; err != nil {
		return err
	}

	return nil
}
