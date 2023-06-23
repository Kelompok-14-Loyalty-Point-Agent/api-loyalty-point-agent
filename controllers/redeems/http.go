package redeems

import (
	"api-loyalty-point-agent/businesses/redeems"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/redeems/request"
	"api-loyalty-point-agent/controllers/redeems/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RedeemsController struct {
	transactionUsecase redeems.Usecase
}

func NewRedeemController(transactionUC redeems.Usecase) *RedeemsController {
	return &RedeemsController{
		transactionUsecase: transactionUC,
	}
}

func (cc *RedeemsController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	redeemData, err := cc.transactionUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	redeems := []response.Redeem{}

	for _, redeem := range redeemData {
		redeems = append(redeems, response.FromDomain(redeem))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all redeem", redeems)
}

func (cc *RedeemsController) RedeemVoucher(c echo.Context) error {
	input := request.Redeem{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	redeem, err := cc.transactionUsecase.RedeemVoucher(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a redeem", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "redeem created", response.FromDomain(redeem))
}
