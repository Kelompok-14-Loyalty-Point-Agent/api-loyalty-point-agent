package request

import (
	"api-loyalty-point-agent/businesses/stock_details"
	"api-loyalty-point-agent/businesses/stocks"
	"api-loyalty-point-agent/businesses/transactions"

	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	Stock          stocks.Domain        `json:"-"`
	Phone          string               `json:"phone" validate:"required"`
	StockDetailsID uint                 `json:"stock_details_id" validate:"required"`
	StockDetails   stock_details.Domain `json:"-"`
	Price          float64              `json:"-"`
	Product        string               `json:"-"`
	Payment_method string               `json:"payment_method" validate:"required"`
	Point          float32              `json:"-"`
	Status         string               `json:"-"`
	Description    string               `json:"-"`
	UserID         uint                 `json:"user_id" validate:"required"`
}

type TransactionPoint struct {
	Point float32 `json:"point" validate:"required"`
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
		Status:         req.Status,
		Description:    req.Description,
		UserID:         req.UserID,
	}
}

func (req *TransactionPoint) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		Point: req.Point,
	}
}

func (req *Transaction) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

func (req *TransactionPoint) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
