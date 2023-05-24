package customers

import (
	"api-loyalty-point-agent/app/middlewares"
	"api-loyalty-point-agent/businesses/customers"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/customers/request"
	"api-loyalty-point-agent/controllers/customers/response"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// AuthController represents the controller for customer authentication.
type AuthController struct {
	authUseCase customers.Usecase
}

// NewAuthController creates a new instance of AuthController.
func NewAuthController(authUC customers.Usecase) *AuthController {
	return &AuthController{
		authUseCase: authUC,
	}
}

func (ctrl *AuthController) GetAllCustomers(c echo.Context) error {
	ctx := c.Request().Context()

	customerData, err := ctrl.authUseCase.GetAllCustomers(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	customers := []response.Customer{}

	for _, customer := range customerData {
		customers = append(customers, response.FromDomain(customer))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all customers", customers)
}

// Register registers a new customer.
// @Summary Register a new customer
// @Description Register a new customer with the given details
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body request.Customer true "example value for registration; email = admin@example.com, name = admin, password = admin123"
// @Success 200 {object} controllers.Response[response.Customer] "success"
// @Success 201 {object} controllers.Response[response.Customer] "success"
// @Failure 400 {object} controllers.Response[string] "failed"
// @Router /customer/register [post]
func (ctrl *AuthController) Register(c echo.Context) error {
	customerInput := request.Customer{}
	ctx := c.Request().Context()

	if err := c.Bind(&customerInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	err := customerInput.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	customer, err := ctrl.authUseCase.Register(ctx, customerInput.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "customer registered", response.FromDomain(customer))
}

// Login logs in a customer.
// @Summary Log in a customer
// @Description Log in a customer with the given email and password
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body request.Customer true "example value for login; email = admin@example.com, name = admin, password = admin123"
// @Success 200 {object} controllers.Response[string] "success"
// @Success 201 {object} controllers.Response[string] "success"
// @Failure 400 {object} controllers.Response[string] "failed"
// @Failure 401 {object} controllers.Response[string] "failed"
// @Router /customer/login [post]
func (ctrl *AuthController) Login(c echo.Context) error {
	customerInput := request.Customer{}
	ctx := c.Request().Context()

	if err := c.Bind(&customerInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	err := customerInput.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	token, err := ctrl.authUseCase.Login(ctx, customerInput.ToDomain())

	var isFailed bool = err != nil || token == ""

	if isFailed {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "token created", token)
}

// Logout logs out a customer.
// @Summary Log out a customer
// @Description Log out a customer by invalidating the token
// @Tags customers
// @Accept json
// @Produce json
// @Success 200 {object} controllers.Response[string] "success"
// @Failure 401 {object} controllers.Response[string] "failed"
// @Router /customer/logout [post]
func (ctrl *AuthController) Logout(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)

	if token == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid token",
		})
	}

	tokenString := token.Raw

	// Invalidate the token by removing it from the whitelist
	isLoggedOut := middlewares.Logout(tokenString)

	if !isLoggedOut {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to logout",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "logout success",
	})
}
