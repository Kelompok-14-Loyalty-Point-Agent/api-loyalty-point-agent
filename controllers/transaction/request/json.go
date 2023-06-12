package request

import (
	"api-loyalty-point-agent/businesses/stocks"
	"api-loyalty-point-agent/businesses/transaction"

	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	// Amount         uint          `json:"amount" validate:"required"`
	Phone          string        `json:"phone" validate:"required"`
	StockID        uint          `json:"stock_id" validate:"required"`
	Stock          stocks.Domain `json:"-"`
	Product        uint          `json:"product" validate:"required"`
	Payment_method uint          `json:"payment_method" validate:"required"`
	Point          uint          `json:"point" validate:"required"`
}

func (req *Transaction) ToDomain() *transaction.Domain {
	return &transaction.Domain{
		// Amount:         req.Amount,
		Phone:          req.Phone,
		StockID:        req.StockID,
		Stock:          req.Stock,
		Product:        req.Product,
		Payment_method: req.Payment_method,
		Point:          req.Point,
	}
}

func (req *Transaction) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
