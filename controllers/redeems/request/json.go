package request

import (
	"api-loyalty-point-agent/businesses/redeems"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Redeem struct {
	Phone     string `json:"phone" validate:"required,NotEmpty"`
	UserID    uint   `json:"user_id" validate:"required"`
	VoucherID uint   `json:"voucher_id" validate:"required"`
}

func (req *Redeem) ToDomain() *redeems.Domain {
	return &redeems.Domain{
		Phone: req.Phone,
		UserID: req.UserID,
		VoucherID: req.VoucherID,
	}
}

func NotEmpty(fl validator.FieldLevel) bool {
	inputData := fl.Field().String()
	inputData = strings.TrimSpace(inputData)

	// Periksa apakah data yang diinput kosong atau hanya berisi spasi
	return inputData != ""
}

func (req *Redeem) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("NotEmpty", NotEmpty)

	err := validate.Struct(req)

	return err
}
