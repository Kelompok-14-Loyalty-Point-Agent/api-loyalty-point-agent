package vouchers

import (
	"api-loyalty-point-agent/businesses/vouchers"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/vouchers/request"
	"api-loyalty-point-agent/controllers/vouchers/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type VoucherController struct {
	transactionUsecase vouchers.Usecase
}

func NewVoucherController(transactionUC vouchers.Usecase) *VoucherController {
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

func (cc *VoucherController) Create(c echo.Context) error {
	input := request.Voucher{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	voucher, err := cc.transactionUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a voucher", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "voucher created", response.FromDomain(voucher))
}
