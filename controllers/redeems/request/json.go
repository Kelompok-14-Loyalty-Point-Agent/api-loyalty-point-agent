package request

import (
	redeems "api-loyalty-point-agent/businesses/redeems"

	// "api-loyalty-point-agent/businesses/stock_details"
	// "api-loyalty-point-agent/businesses/transactions"

	"time"

	"github.com/go-playground/validator/v10"
)

type Redeem struct {
	CreatedAt time.Time `json:"created_at"`
	Phone     string    `json:"phone"`
	Cost      uint      `json:"cost"`
}

func (req *Redeem) ToDomain() *redeems.Domain {
	return &redeems.Domain{
		CreatedAt: req.CreatedAt,
		Phone:     req.Phone,
		Cost:      req.Cost,
	}
}

func (req *Redeem) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}