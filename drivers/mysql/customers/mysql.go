package customers

import (
	"context"
	"api-loyalty-point-agent/businesses/customers"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type customerRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) customers.Repository {
	return &customerRepository{
		conn: conn,
	}
}

func (ur *customerRepository) Register(ctx context.Context, customerDomain *customers.Domain) (customers.Domain, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(customerDomain.Password), bcrypt.DefaultCost)

	if err != nil {
		return customers.Domain{}, err
	}

	record := FromDomain(customerDomain)

	record.Password = string(password)

	result := ur.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return customers.Domain{}, err
	}

	err = result.Last(&record).Error

	if err != nil {
		return customers.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (ur *customerRepository) GetByEmail(ctx context.Context, customerDomain *customers.Domain) (customers.Domain, error) {
	var customer Customer

	err := ur.conn.WithContext(ctx).First(&customer, "email = ?", customerDomain.Email).Error

	if err != nil {
		return customers.Domain{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(customerDomain.Password))

	if err != nil {
		return customers.Domain{}, err
	}

	return customer.ToDomain(), nil
}