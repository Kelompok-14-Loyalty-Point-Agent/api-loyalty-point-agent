package vouchers

import (
	"api-loyalty-point-agent/businesses/vouchers"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/vouchers/response"
	_redeemRequest "api-loyalty-point-agent/controllers/redeems/request"
	_redeemResponse "api-loyalty-point-agent/controllers/redeems/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type VoucherController struct {
	voucherUsecase vouchers.Usecase
}

func NewVoucherController(voucherUC vouchers.Usecase) *VoucherController {
	return &VoucherController{
		voucherUsecase: voucherUC,
	}
}

func (cc *VoucherController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	voucherData, err := cc.voucherUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	vouchers := []response.Voucher{}

	for _, voucher := range voucherData {
		vouchers = append(vouchers, response.FromDomain(voucher))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all voucher", vouchers)
}

func (ctrl *VoucherController) GetByID(c echo.Context) error {
	var voucherID string = c.Param("id")

	ctx := c.Request().Context()

	user, err := ctrl.voucherUsecase.GetByID(ctx, voucherID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "get voucher by id", response.FromDomain(user))
}

func (cc *VoucherController) RedeemVoucher(c echo.Context) error {
	input := _redeemRequest.Redeem{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	redeem, err := cc.voucherUsecase.RedeemVoucher(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "redeem created", _redeemResponse.FromDomain(redeem))
}
