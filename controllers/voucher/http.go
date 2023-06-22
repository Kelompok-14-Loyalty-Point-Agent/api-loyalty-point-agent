package voucher

import (
	"api-loyalty-point-agent/businesses/voucher"
	"api-loyalty-point-agent/controllers"

	// "api-loyalty-point-agent/controllers/voucher/response"
	"api-loyalty-point-agent/controllers/voucher/request"

	// "api-loyalty-point-agent/controllers/transactions/request"
	"api-loyalty-point-agent/controllers/voucher/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type VoucherController struct {
	transactionUsecase voucher.Usecase
}

func NewVoucherController(transactionUC voucher.Usecase) *VoucherController {
	return &VoucherController{
		transactionUsecase: transactionUC,
	}
}

func (cc *VoucherController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	voucherData, err := cc.transactionUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	vouchers := []response.Voucher{}

	for _, voucher := range voucherData {
		vouchers = append(vouchers, response.FromDomain(voucher))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all voucher", vouchers)
}

func (cc *VoucherController) RedeemVoucher(c echo.Context) error {
	input := request.Voucher{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	voucher, err := cc.transactionUsecase.RedeemVoucher(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a voucher", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "voucher created", response.FromDomain(voucher))
}
