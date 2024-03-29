package redeems

import (
	"api-loyalty-point-agent/businesses/redeems"
	"api-loyalty-point-agent/controllers"
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

func (ctrl *RedeemsController) GetByID(c echo.Context) error {
	var redeemID string = c.Param("id")

	ctx := c.Request().Context()

	user, err := ctrl.transactionUsecase.GetByID(ctx, redeemID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "get redeem by id", response.FromDomain(user))
}