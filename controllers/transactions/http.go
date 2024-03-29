package transactions

import (
	"api-loyalty-point-agent/businesses/transactions"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/transactions/request"
	"api-loyalty-point-agent/controllers/transactions/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	transactionUsecase transactions.Usecase
}

func NewTransactionController(transactionUC transactions.Usecase) *TransactionController {
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

	transactions := []response.TransactionInAdmin{}

	for _, transaction := range transactionData {
		transactions = append(transactions, response.FromDomainInAdmin(transaction))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all transactions", transactions)
}

func (cc *TransactionController) GetByID(c echo.Context) error {
	var transactionID string = c.Param("id")
	ctx := c.Request().Context()

	transaction, err := cc.transactionUsecase.GetByID(ctx, transactionID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "transaction not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "transaction found", response.FromDomainInAdmin(transaction))
}

func (cc *TransactionController) Create(c echo.Context) error {
	input := request.Transaction{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	transaction, err := cc.transactionUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "transaction created", response.FromDomain(transaction))
}

func (cc *TransactionController) GetAllByUserID(c echo.Context) error {
	var userID string = c.Param("id")
	ctx := c.Request().Context()

	transactionData, err := cc.transactionUsecase.GetAllByUserID(ctx, userID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", err.Error(), "")
	}

	transactions := []response.Transaction{}

	for _, transaction := range transactionData {
		transactions = append(transactions, response.FromDomain(transaction))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all transactions of an user", transactions)
}

func (cc *TransactionController) GetAllByUserIDSorted(c echo.Context) error {
	var userID string = c.Param("id")
	ctx := c.Request().Context()

	transactionData, err := cc.transactionUsecase.GetAllByUserIDSorted(ctx, userID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", err.Error(), "")
	}

	transactions := []response.Transaction{}

	for _, transaction := range transactionData {
		transactions = append(transactions, response.FromDomain(transaction))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all transactions of an user sorted by recent transaction", transactions)
}

func (cc *TransactionController) UpdatePoint(c echo.Context) error {
	var transactionID string = c.Param("id")
	input := request.TransactionPoint{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	transaction, err := cc.transactionUsecase.UpdatePoint(ctx, input.ToDomain(), transactionID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "transaction point updated", response.FromDomainInAdmin(transaction))
}
