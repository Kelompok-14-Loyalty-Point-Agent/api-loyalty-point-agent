package request

import (
	"api-loyalty-point-agent/businesses/providers"

	"github.com/go-playground/validator/v10"
)

type Provider struct {
	Name string `json:"name" validate:"required"`
	URL  string `json:"url" validate:"required"`
}

func (req *Provider) ToDomain() *providers.Domain {
	return &providers.Domain{
		Name: req.Name,
		URL:  req.URL,
	}
}

func (req *Provider) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
