package request

import (
	"api-loyalty-point-agent/businesses/user_details"

	"github.com/go-playground/validator/v10"
)

type UserDetail struct {
	Member           string `json:"member"`
	TPoint           int64  `json:"tPoint"`
	Age              string `json:"age" `
	Gender           string `json:"gender"`
	Address          string `json:"address"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phone_number"`
	URL              string `json:"url"`
	CountTransaction int64  `json:"count_transaction"`
	CountRedeem      int64  `json:"count_redeem"`
}

func (req *UserDetail) ToDomain() *user_details.Domain {
	return &user_details.Domain{
		Member:           req.Member,
		TPoint:           req.TPoint,
		Age:              req.Age,
		Gender:           req.Gender,
		Address:          req.Address,
		PhoneNumber:      req.PhoneNumber,
		URL:              req.URL,
		CountTransaction: req.CountTransaction,
		CountRedeem:      req.CountRedeem,
	}
}

func (req *UserDetail) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
