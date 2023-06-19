package users

import (
	"api-loyalty-point-agent/app/middlewares"
	"api-loyalty-point-agent/businesses/users"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/users/request"
	"api-loyalty-point-agent/controllers/users/response"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// AuthController represents the controller for user authentication.
type AuthController struct {
	authUseCase users.Usecase
}

// NewAuthController creates a new instance of AuthController.
func NewAuthController(authUC users.Usecase) *AuthController {
	return &AuthController{
		authUseCase: authUC,
	}
}

// Get Get all users.
// @Summary Retrieves all users data
// @Description Retrieves all users data with the given details
// @Tags users
// @Accept json
// @Produce json
// @Security BearerToken
// @Success 200 {object} controllers.Response[response.User] "success"
// @Success 201 {object} controllers.Response[response.User] "success"
// @Failure 400 {object} controllers.Response[string] "failed"
// @Router /users/usersAll [get]
func (ctrl *AuthController) GetAllCustomers(c echo.Context) error {
	ctx := c.Request().Context()

	userData, err := ctrl.authUseCase.GetAllCustomers(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	users := []response.User{}

	for _, user := range userData {
		users = append(users, response.FromDomain(user))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all customers", users)
}

// Register registers a new user.
// @Summary Register a new user
// @Description Register a new user with the given details
// @Tags users
// @Accept json
// @Produce json
// @Param user body request.User true "example value for registration; email = admin@example.com, name = admin, password = admin123"
// @Success 200 {object} controllers.Response[response.User] "success"
// @Success 201 {object} controllers.Response[response.User] "success"
// @Failure 400 {object} controllers.Response[string] "failed"
// @Router /user/register [post]
func (ctrl *AuthController) Register(c echo.Context) error {
	userInput := request.UserRegistration{}
	ctx := c.Request().Context()

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	err := userInput.ValidateRegistration()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	user, err := ctrl.authUseCase.Register(ctx, userInput.ToDomainRegistration())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "user registered", response.FromDomain(user))
}

// Login logs in a user.
// @Summary Log in a user
// @Description Log in a user with the given email and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body request.User true "example value for login; email = admin@example.com, name = admin, password = admin123"
// @Success 200 {object} controllers.Response[string] "success"
// @Success 201 {object} controllers.Response[string] "success"
// @Failure 400 {object} controllers.Response[string] "failed"
// @Failure 401 {object} controllers.Response[string] "failed"
// @Router /user/login [post]
func (ctrl *AuthController) Login(c echo.Context) error {
	userInput := request.UserLogin{}
	ctx := c.Request().Context()

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	err := userInput.ValidateLogin()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	token, err := ctrl.authUseCase.Login(ctx, userInput.ToDomainLogin())

	var isFailed bool = err != nil || token == ""

	if isFailed {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "token created", token)
}

// Logout logs out a user.
// @Summary Log out a user
// @Description Log out a user by invalidating the token
// @Tags users
// @Accept json
// @Produce json
// @Security BearerToken
// @Success 200 {object} controllers.Response[string] "success"
// @Failure 401 {object} controllers.Response[string] "failed"
// @Router /users/logout [post]
func (ctrl *AuthController) Logout(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)

	isListed := middlewares.CheckToken(token.Raw)

	if !isListed {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid token", isListed)
	}
	// Invalidate the token by removing it from the whitelist
	isLoggedOut := middlewares.Logout(token.Raw)

	return controllers.NewResponse(c, http.StatusOK, "success", "logout success", isLoggedOut)
}

func (ctrl *AuthController) GetByID(c echo.Context) error {
	var userID string = c.Param("id")

	ctx := c.Request().Context()

	user, err := ctrl.authUseCase.GetByID(ctx, userID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "get user by id", response.FromDomain(user))
}

func (ctrl *AuthController) UpdateProfileCustomer(c echo.Context) error {
	var userID string = c.Param("id")
	input := request.CustomerProfile{}

	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	user, err := ctrl.authUseCase.UpdateProfileCustomer(ctx, input.ToDomainProfileCustomer(), userID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "customer updated", response.FromDomain(user))
}

func (ctrl *AuthController) UpdateProfileAdmin(c echo.Context) error {
	var userID string = c.Param("id")
	input := request.AdminProfile{}

	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	user, err := ctrl.authUseCase.UpdateProfileAdmin(ctx, input.ToDomainProfileAdmin(), userID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "admin updated", response.FromDomain(user))
}

func (ctrl *AuthController) ChangePassword(c echo.Context) error {
	var userID string = c.Param("id")
	input := request.InputPassword{}

	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	user, err := ctrl.authUseCase.ChangePassword(ctx, input.ToDomainProfilePassword(), userID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "password changed", response.FromDomain(user))
}

func (ctrl *AuthController) DeleteCustomer(c echo.Context) error {
	var userID string = c.Param("id")

	ctx := c.Request().Context()

	err := ctrl.authUseCase.DeleteCustomer(ctx, userID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "customer deleted", "")
}

