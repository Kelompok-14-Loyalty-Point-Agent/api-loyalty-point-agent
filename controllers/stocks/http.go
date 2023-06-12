package stocks

import (
	"api-loyalty-point-agent/businesses/stocks"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/stocks/request"
	"api-loyalty-point-agent/controllers/stocks/response"
	_reqStockTransaction "api-loyalty-point-agent/controllers/stock_transactions/request"
	_resStockTransaction "api-loyalty-point-agent/controllers/stock_transactions/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StockController struct {
	stockUsecase stocks.Usecase
}

func NewStockController(stockUC stocks.Usecase) *StockController {
	return &StockController{
		stockUsecase: stockUC,
	}
}

func (cc *StockController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	stocksData, err := cc.stockUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	stocks := []response.Stock{}

	for _, stock := range stocksData {
		stocks = append(stocks, response.FromDomain(stock))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all stocks", stocks)
}

func (cc *StockController) GetByID(c echo.Context) error {
	var stockID string = c.Param("id")
	ctx := c.Request().Context()

	stock, err := cc.stockUsecase.GetByID(ctx, stockID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "stock not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "stock found", response.FromDomain(stock))
}

func (cc *StockController) AddStock(c echo.Context) error {
	input := _reqStockTransaction.StockTransaction{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	stock_transaction, err := cc.stockUsecase.AddStock(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "stock transaction created", _resStockTransaction.FromDomain(stock_transaction))
}

func (cc *StockController) Update(c echo.Context) error {
	var stockID string = c.Param("id")
	ctx := c.Request().Context()

	input := request.Stock{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	stock, err := cc.stockUsecase.Update(ctx, input.ToDomain(), stockID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update stock failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "stock updated", response.FromDomain(stock))
}

func (cc *StockController) Delete(c echo.Context) error {
	var stockID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.stockUsecase.Delete(ctx, stockID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete stock failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "stock deleted", "")
}