package customers

import (
	"api-loyalty-point-agent/businesses/customers"
	"context"

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

func (ur *customerRepository) GetAllCustomers(ctx context.Context) ([]customers.Domain, error) {
	var records []Customer

	if err := ur.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	customers := []customers.Domain{}

	for _, customer := range records {
		customers = append(customers, customer.ToDomain())
	}

	return customers, nil
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
