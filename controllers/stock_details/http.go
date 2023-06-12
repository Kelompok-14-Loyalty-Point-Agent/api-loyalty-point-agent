package stock_details

import (
	"api-loyalty-point-agent/businesses/stock_details"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/stock_details/request"
	"api-loyalty-point-agent/controllers/stock_details/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StockDetailController struct {
	stock_detailUsecase stock_details.Usecase
}

func NewStockDetailController(stock_detailUC stock_details.Usecase) *StockDetailController {
	return &StockDetailController{
		stock_detailUsecase: stock_detailUC,
	}
}

func (cc *StockDetailController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	stock_detailsData, err := cc.stock_detailUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	stock_details := []response.StockDetail{}

	for _, stock_detail := range stock_detailsData {
		stock_details = append(stock_details, response.FromDomain(stock_detail))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all stock_details", stock_details)
}

func (cc *StockDetailController) GetByID(c echo.Context) error {
	var stock_detailID string = c.Param("id")
	ctx := c.Request().Context()

	stock_detail, err := cc.stock_detailUsecase.GetByID(ctx, stock_detailID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "stock_detail not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "stock_detail found", response.FromDomain(stock_detail))
}

func (cc *StockDetailController) Create(c echo.Context) error {
	input := request.StockDetail{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	if input.StockID == 0 {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "stock_id is required", "")
	}

	stock_detail, err := cc.stock_detailUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a stock_detail", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "stock_detail created", response.FromDomain(stock_detail))
}

func (cc *StockDetailController) Update(c echo.Context) error {
	var stock_detailID string = c.Param("id")
	ctx := c.Request().Context()

	input := request.StockDetail{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	stock_detail, err := cc.stock_detailUsecase.Update(ctx, input.ToDomain(), stock_detailID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update stock_detail failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "stock_detail updated", response.FromDomain(stock_detail))
}

func (cc *StockDetailController) Delete(c echo.Context) error {
	var stock_detailID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.stock_detailUsecase.Delete(ctx, stock_detailID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete stock_detail failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "stock_detail deleted", "")
}