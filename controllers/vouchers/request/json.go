package request

import (
	voucher "api-loyalty-point-agent/businesses/vouchers"

	// "api-loyalty-point-agent/businesses/stock_details"
	// "api-loyalty-point-agent/businesses/transactions"

	"github.com/go-playground/validator/v10"
)

type Voucher struct {
	// ID string `json:"id"`
	Product string `json:"product"`
	Benefit string `json:"benefit"`
	Cost    float32   `json:"cost"`
}

func (req *Voucher) ToDomain() *voucher.Domain {
	return &voucher.Domain{
		Product: req.Product,
		Benefit: req.Benefit,
		Cost:    req.Cost,
	}
}

func (req *Voucher) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
