package request

import (
	"api-loyalty-point-agent/businesses/users"
	"strings"

	"github.com/go-playground/validator/v10"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}

type UserRegistration struct {
	Name     string `json:"name" validate:"required,nameNotEmpty"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}

func (req *UserLogin) ToDomainLogin() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
}

func (req *UserRegistration) ToDomainRegistration() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
}

func (req *UserLogin) ValidateLogin() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func NotEmpty(fl validator.FieldLevel) bool {
	inputData := fl.Field().String()
	inputData = strings.TrimSpace(inputData)

	// Periksa apakah data yang diinput kosong atau hanya berisi spasi
	return inputData != ""
}

func (req *UserRegistration) ValidateRegistration() error {
	validate := validator.New()
	validate.RegisterValidation("nameNotEmpty", NotEmpty)

	err := validate.Struct(req)

	return err
}
