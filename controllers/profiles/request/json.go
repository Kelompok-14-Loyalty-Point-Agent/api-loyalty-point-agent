package request

import (
	"api-loyalty-point-agent/businesses/profiles"

	"github.com/go-playground/validator/v10"
)

type UpdateProfileRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type ChangePasswordRequest struct {
	Password string `json:"password" validate:"required"`
}

func (req *UpdateProfileRequest) ToDomain() *profiles.Domain {
	return &profiles.Domain{
		Name:    req.Name,
		Address: req.Address,
	}
}

func (req *UpdateProfileRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func (req *ChangePasswordRequest) ToDomain() *profiles.Domain {
	return &profiles.Domain{
		Password: req.Password,
	}
}

func (req *ChangePasswordRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
