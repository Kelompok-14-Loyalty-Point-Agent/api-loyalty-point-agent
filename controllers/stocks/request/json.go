package request

import (
	"api-loyalty-point-agent/businesses/providers"
	"api-loyalty-point-agent/businesses/stocks"

	"github.com/go-playground/validator/v10"
)

type Stock struct {
	Name       string           `json:"name"`
	Type       string           `json:"type"`
	Stock      float64          `json:"stock"`
	Price      float64          `json:"price"`
	Quantity   float64          `json:"quantity"`
	ProviderID uint             `json:"provider_id"`
	Provider   providers.Domain `json:"-"`
}

func (req *Stock) ToDomain() *stocks.Domain {
	return &stocks.Domain{
		Name:       req.Name,
		Type:       req.Type,
		Stock:      req.Stock,
		Price:      req.Price,
		Quantity:   req.Quantity,
		Provider:   req.Provider,
		ProviderID: req.ProviderID,
	}
}

func (req *Stock) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
