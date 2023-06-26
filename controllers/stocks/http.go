package stocks

import (
	"api-loyalty-point-agent/app/middlewares"
	"api-loyalty-point-agent/businesses/stocks"
	"api-loyalty-point-agent/controllers"
	_reqStockTransaction "api-loyalty-point-agent/controllers/stock_transactions/request"
	_resStockTransaction "api-loyalty-point-agent/controllers/stock_transactions/response"
	"api-loyalty-point-agent/controllers/stocks/response"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
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
	token := c.Get("user").(*jwt.Token)

	isListed := middlewares.CheckToken(token.Raw)

	if !isListed {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid token", isListed)
	}

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
	token := c.Get("user").(*jwt.Token)

	isListed := middlewares.CheckToken(token.Raw)

	if !isListed {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid token", isListed)
	}

	stock, err := cc.stockUsecase.GetByID(ctx, stockID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "stock not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "stock found", response.FromDomain(stock))
}

func (cc *StockController) AddStock(c echo.Context) error {
	input := _reqStockTransaction.StockTransaction{}
	ctx := c.Request().Context()
	token := c.Get("user").(*jwt.Token)

	isListed := middlewares.CheckToken(token.Raw)

	if !isListed {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid token", isListed)
	}

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
