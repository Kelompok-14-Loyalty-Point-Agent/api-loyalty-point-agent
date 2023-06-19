package request

import (
	"api-loyalty-point-agent/businesses/profiles"
	"api-loyalty-point-agent/businesses/users"
	"strings"

	"github.com/go-playground/validator/v10"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}

// user registration with Name
type UserRegistration struct {
	Name     string `json:"name" validate:"required,nameNotEmpty"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}

type CustomerProfile struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Age     uint   `json:"age"`
	Gender  string `json:"gender"`
	Phone   string `json:"phone"`
}

type AdminProfile struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type InputPassword struct {
	Password string `json:"password" validate:"required"`
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

func (req *CustomerProfile) ToDomainProfileCustomer() *users.Domain {
	return &users.Domain{
		Name:  req.Name,
		Email: req.Email,
		Profile: profiles.Domain{
			Address: req.Address,
			Phone:   req.Phone,
			Age:     req.Age,
			Gender:  req.Gender,
		},
	}
}

func (req *AdminProfile) ToDomainProfileAdmin() *users.Domain {
	return &users.Domain{
		Name:  req.Name,
		Profile: profiles.Domain{
			Address: req.Address,
		},
	}
}

func (req *InputPassword) ToDomainProfilePassword() *users.Domain {
	return &users.Domain{
		Password: req.Password,
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
