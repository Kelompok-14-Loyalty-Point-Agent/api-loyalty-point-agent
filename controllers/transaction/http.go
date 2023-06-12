package transaction

import (
	"api-loyalty-point-agent/businesses/transaction"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/transaction/request"
	"api-loyalty-point-agent/controllers/transaction/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	transactionUsecase transaction.Usecase
}

func NewTransactionController(transactionUC transaction.Usecase) *TransactionController {
	return &TransactionController{
		transactionUsecase: transactionUC,
	}
}

func (cc *TransactionController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	transactionData, err := cc.transactionUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	transactions := []response.Transaction{}

	for _, transaction := range transactionData {
		transactions = append(transactions, response.FromDomain(transaction))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all transactions", transactions)
}

func (cc *TransactionController) GetByID(c echo.Context) error {
	var transactionID string = c.Param("id")
	ctx := c.Request().Context()

	transaction, err := cc.transactionUsecase.GetByID(ctx, transactionID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "stock_detail not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "stock_detail found", response.FromDomain(transaction))
}

func (cc *TransactionController) Create(c echo.Context) error {
	input := request.Transaction{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	transaction, err := cc.transactionUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a stock_detail", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "stock_detail created", response.FromDomain(transaction))
}
