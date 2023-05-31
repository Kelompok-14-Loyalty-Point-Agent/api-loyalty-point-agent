package request

import (
	"api-loyalty-point-agent/businesses/providers"
	"api-loyalty-point-agent/businesses/stocks"

	"github.com/go-playground/validator/v10"
)

type Stock struct {
	Type       string           `json:"type" validate:"required"`
	TotalStock float64          `json:"total_stock" validate:"required"`
	ProviderID uint             `json:"provider_id" validate:"required"`
	Provider   providers.Domain `json:"-"`
}

func (req *Stock) ToDomain() *stocks.Domain {
	return &stocks.Domain{
		Type:       req.Type,
		TotalStock: req.TotalStock,
		Provider:   req.Provider,
		ProviderID: req.ProviderID,
	}
}

func (req *Stock) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
