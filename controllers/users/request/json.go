package request

import (
	"api-loyalty-point-agent/businesses/users"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}

func (req *User) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
}

func (req *User) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
