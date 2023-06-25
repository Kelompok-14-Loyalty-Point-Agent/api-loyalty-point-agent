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

type UserRegistration struct {
	Name     string `json:"name" validate:"required,NotEmpty"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}

type CustomerProfile struct {
	Name    string `json:"name" validate:"NotEmpty"`
	Email   string `json:"email" validate:"email"`
	Address string `json:"address" validate:"NotEmpty"`
	Age     uint   `json:"age" validate:"NotEmpty"`
	Gender  string `json:"gender" validate:"NotEmpty"`
	Phone   string `json:"phone" validate:"NotEmpty"`
}

type AdminProfile struct {
	Name    string `json:"name" validate:"NotEmpty"`
	Address string `json:"address" validate:"NotEmpty"`
}

type InputPassword struct {
	Password string `json:"password" validate:"required"`
}

type CustomerProfileInAdmin struct {
	Email string `json:"email" validate:"NotEmpty"`
	Phone string `json:"phone" validate:"NotEmpty"`
}

func (req *UserLogin) ToDomainLogin() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
}

func (req *UserRegistration) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
}

func (req *CustomerProfile) ToDomain() *users.Domain {
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

func (req *AdminProfile) ToDomain() *users.Domain {
	return &users.Domain{
		Name: req.Name,
		Profile: profiles.Domain{
			Address: req.Address,
		},
	}
}

func (req *InputPassword) ToDomain() *users.Domain {
	return &users.Domain{
		Password: req.Password,
	}
}

func (req *CustomerProfileInAdmin) ToDomain() *users.Domain {
	return &users.Domain{
		Email: req.Email,
		Profile: profiles.Domain{
			Phone: req.Phone,
		},
	}
}

func validateRequest(req interface{}) error {
	validate := validator.New()
	validate.RegisterValidation("NotEmpty", NotEmpty)

	err := validate.Struct(req)

	return err
}

func NotEmpty(fl validator.FieldLevel) bool {
	inputData := fl.Field().String()
	inputData = strings.TrimSpace(inputData)

	return inputData != ""
}

func (req *UserLogin) Validate() error {
	return validateRequest(req)
}

func (req *UserRegistration) Validate() error {
	return validateRequest(req)
}

func (req *CustomerProfile) Validate() error {
	return validateRequest(req)
}

func (req *AdminProfile) Validate() error {
	return validateRequest(req)
}

func (req *InputPassword) Validate() error {
	return validateRequest(req)
}

func (req *CustomerProfileInAdmin) Validate() error {
	return validateRequest(req)
}
