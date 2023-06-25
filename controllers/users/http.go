package users

import (
	"api-loyalty-point-agent/app/middlewares"
	"api-loyalty-point-agent/businesses/users"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/users/request"
	"api-loyalty-point-agent/controllers/users/response"
	"api-loyalty-point-agent/drivers/aws"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authUseCase users.Usecase
}

func NewAuthController(authUC users.Usecase) *AuthController {
	return &AuthController{
		authUseCase: authUC,
	}
}

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

func (ctrl *AuthController) Register(c echo.Context) error {
	userInput := request.UserRegistration{}
	ctx := c.Request().Context()

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	err := userInput.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	user, err := ctrl.authUseCase.Register(ctx, userInput.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "user registered", response.FromDomain(user))
}

func (ctrl *AuthController) Login(c echo.Context) error {
	userInput := request.UserLogin{}
	ctx := c.Request().Context()

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	err := userInput.Validate()

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

	err := input.Validate()
	if err!= nil{
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	user, err := ctrl.authUseCase.UpdateProfileCustomer(ctx, input.ToDomain(), userID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "customer updated", response.FromDomain(user))
}

func (ctrl *AuthController) UpdateProfileCustomerInAdmin(c echo.Context) error {
	var userID string = c.Param("id")
	input := request.CustomerProfileInAdmin{}

	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := input.Validate()
	if err!= nil{
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	user, err := ctrl.authUseCase.UpdateProfileCustomerInAdmin(ctx, input.ToDomain(), userID)

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

	err := input.Validate()
	if err!= nil{
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	user, err := ctrl.authUseCase.UpdateProfileAdmin(ctx, input.ToDomain(), userID)

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

	err := input.Validate()
	if err!= nil{
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	user, err := ctrl.authUseCase.ChangePassword(ctx, input.ToDomain(), userID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "password changed", response.FromDomain(user))
}

func (ctrl *AuthController) ChangePicture(c echo.Context) error {
	var userID string = c.Param("id")
	file, err := c.FormFile("filename")

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "failed to upload file", "")
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	ctx := c.Request().Context()

	url, err := aws_driver.UploadFileToBucket("user.png", src)

	_, prev_url, err := ctrl.authUseCase.ChangePicture(ctx, url, userID)

	if err != nil {
		err = aws_driver.DeleteFileFromBucket(url)
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	err = aws_driver.DownloadFileFromBucket(url, "./assets/users/")

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	if len(prev_url) != 0 {
		local_url := "./assets/users/" + prev_url

		err = os.Remove(local_url)

		if err != nil {
			return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
		}

		err = aws_driver.DeleteFileFromBucket(prev_url)
		if err != nil {
			return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
		}
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "change picture successfully", url)
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


