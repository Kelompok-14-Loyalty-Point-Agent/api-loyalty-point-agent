package stock_transactions

import (
	"api-loyalty-point-agent/app/middlewares"
	"api-loyalty-point-agent/businesses/stock_transactions"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/stock_transactions/response"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type StockTransactionController struct {
	stock_transactionUsecase stock_transactions.Usecase
}

func NewStockTransactionController(stock_transactionUC stock_transactions.Usecase) *StockTransactionController {
	return &StockTransactionController{
		stock_transactionUsecase: stock_transactionUC,
	}
}

func (cc *StockTransactionController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	token := c.Get("user").(*jwt.Token)

	isListed := middlewares.CheckToken(token.Raw)

	if !isListed {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid token", isListed)
	}

	stock_transactionsData, err := cc.stock_transactionUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	stock_transactions := []response.StockTransaction{}

	for _, stock_transaction := range stock_transactionsData {
		stock_transactions = append(stock_transactions, response.FromDomain(stock_transaction))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all stock_transactions", stock_transactions)
}

func (cc *StockTransactionController) GetByID(c echo.Context) error {
	var stock_transactionID string = c.Param("id")
	ctx := c.Request().Context()
	token := c.Get("user").(*jwt.Token)

	isListed := middlewares.CheckToken(token.Raw)

	if !isListed {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid token", isListed)
	}

	stock_transaction, err := cc.stock_transactionUsecase.GetByID(ctx, stock_transactionID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "stock_transaction not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "stock_transaction found", response.FromDomain(stock_transaction))
}
