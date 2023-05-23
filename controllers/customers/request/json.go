package request

import (
	"api-loyalty-point-agent/businesses/customers"

	"github.com/go-playground/validator/v10"
)

type Customer struct {
	Name      	string  `json:"name" validate:"required"`
	Email    	string 	`json:"email" validate:"required,email"`
	Password 	string 	`json:"password" validate:"required"`
}

func (req *Customer) ToDomain() *customers.Domain {
	return &customers.Domain{
		Name: 		req.Name,
		Email:    	req.Email,
		Password: 	req.Password,
	}
}

func (req *Customer) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}