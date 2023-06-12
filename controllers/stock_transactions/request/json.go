package request

import (
	"api-loyalty-point-agent/businesses/stock_transactions"

	"github.com/go-playground/validator/v10"
)

type StockTransaction struct {
	InputStock    float64 `json:"input_stock" validate:"required"`
	PayAmount     float64 `json:"-"`
	PaymentMethod string  `json:"payment_method" validate:"required"`
	ProviderName  string  `json:"provider_name" validate:"required"`
	Status        string  `json:"-"`
	StockID       uint    `json:"stock_id" validate:"required"`
	UserID        uint    `json:"user_id" validate:"required"`
}

func (req *StockTransaction) ToDomain() *stock_transactions.Domain {
	return &stock_transactions.Domain{
		InputStock:    req.InputStock,
		PayAmount:     req.PayAmount,
		PaymentMethod: req.PaymentMethod,
		ProviderName:  req.ProviderName,
		Status:        req.Status,
		StockID:       req.StockID,
		UserID:        req.UserID,
	}
}

func (req *StockTransaction) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
