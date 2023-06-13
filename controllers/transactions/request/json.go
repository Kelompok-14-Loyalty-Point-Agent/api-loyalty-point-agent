package request

import (
	"api-loyalty-point-agent/businesses/stocks"

	"api-loyalty-point-agent/businesses/stock_details"
	"api-loyalty-point-agent/businesses/transactions"

	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	Stock          stocks.Domain        `json:"-"`
	Phone          string               `json:"phone" validate:"required"`
	StockDetailsID uint                 `json:"stock_details_id" validate:"required"`
	StockDetails   stock_details.Domain `json:"-"`
	Price          float64              `json:"price"`
	Product        uint                 `json:"product" validate:"required"`
	Payment_method uint                 `json:"payment_method" validate:"required"`
	Point          uint                 `json:"point" validate:"required"`
}

func (req *Transaction) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		Phone:          req.Phone,
		StockDetailsID: req.StockDetailsID,
		StockDetails:   req.StockDetails,
		StockID:        req.StockDetails.StockID,
		Price:          req.Price,
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
