package request

import (
	"api-loyalty-point-agent/businesses/stock_details"

	"github.com/go-playground/validator/v10"
)

type StockDetail struct {
	StockID  uint    `json:"stock_id" validate:"required"`
	Stock    float64 `json:"stock" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	Quantity float64 `json:"quantity" validate:"required"`
}

func (req *StockDetail) ToDomain() *stock_details.Domain {
	return &stock_details.Domain{
		StockID:     req.StockID,
		Stock:       req.Stock,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}
}

func (req *StockDetail) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
