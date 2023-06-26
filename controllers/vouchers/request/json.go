package request

import (
	"api-loyalty-point-agent/businesses/vouchers"

	"github.com/go-playground/validator/v10"
)

type Voucher struct {
	Product string  `json:"product"`
	Benefit string  `json:"benefit"`
	Cost    float32 `json:"cost"`
}

func (req *Voucher) ToDomain() *vouchers.Domain {
	return &vouchers.Domain{
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
