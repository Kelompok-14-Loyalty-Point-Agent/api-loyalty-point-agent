package request

import (
	"api-loyalty-point-agent/businesses/voucher"

	// "api-loyalty-point-agent/businesses/stock_details"
	// "api-loyalty-point-agent/businesses/transactions"

	"github.com/go-playground/validator/v10"
)

type Voucher struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Cost  uint   `json:"cost"`
	Point uint   `json:"point"`
	// Expired	uint
	// Date_Exhange time.Time `json:"date_exhange"`
}

func (req *Voucher) ToDomain() *voucher.Domain {
	return &voucher.Domain{
		Title: req.Title,
		Url:   req.Url,
		Cost:  req.Cost,
		Point: req.Point,
		// Status:         domain.Status,
		// Date_Exhange: req.Date_Exhange,
	}
}

func (req *Voucher) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
