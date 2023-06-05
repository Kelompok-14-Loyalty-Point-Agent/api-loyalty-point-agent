package user_details

import (
	"api-loyalty-point-agent/businesses/user_details"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/user_details/request"
	"api-loyalty-point-agent/controllers/user_details/response"
	"encoding/json"

	// "io"

	aws_driver "api-loyalty-point-agent/drivers/aws"
	// "api-loyalty-point-agent/utils"
	"net/http"
	// "os"
	// "path/filepath"

	"github.com/labstack/echo/v4"
)

type UserDetailController struct {
	user_detailUsecase user_details.Usecase
}

func NewUserDetailController(user_detailUC user_details.Usecase) *UserDetailController {
	return &UserDetailController{
		user_detailUsecase: user_detailUC,
	}
}

func (cc *UserDetailController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	user_detailsData, err := cc.user_detailUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	user_details := []response.UserDetail{}

	for _, user_detail := range user_detailsData {
		user_details = append(user_details, response.FromDomain(user_detail))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all user_details", user_details)
}

func (cc *UserDetailController) GetByID(c echo.Context) error {
	var user_detailID string = c.Param("id")
	ctx := c.Request().Context()

	user_detail, err := cc.user_detailUsecase.GetByID(ctx, user_detailID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "user_detail not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "user_detail found", response.FromDomain(user_detail))
}

func (cc *UserDetailController) ReadFile(c echo.Context) error {
	key, err := aws_driver.ReadAllFilesFromBucket()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	_, err = json.Marshal(&key)

	if err != nil {
		return err
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "read data from bucket", key)
}

func (cc *UserDetailController) DownloadFile(c echo.Context) error {
	name := c.FormValue("filename")

	err := aws_driver.DownloadFileFromBucket(name, "./assets/user_details/")

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "downloaded data from bucket", "")
}

func (cc *UserDetailController) Update(c echo.Context) error {
	var user_detailID string = c.Param("id")

	ctx := c.Request().Context()

	input := request.UserDetail{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	user_detail, err := cc.user_detailUsecase.Update(ctx, input.ToDomain(), user_detailID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update user_detail failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "user_detail updated", response.FromDomain(user_detail))
}

func (cc *UserDetailController) UpdatePicture(c echo.Context) error {
	var user_detailID string = c.Param("id")

	ctx := c.Request().Context()

	file, err := c.FormFile("picture")

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "failed to upload file", "")
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	result, err := aws_driver.UploadFileToBucket(file.Filename, src)
	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	input := request.UserDetail{}

	input.URL = result

	user_detail, err := cc.user_detailUsecase.Update(ctx, input.ToDomain(), user_detailID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update user_detail failed", "")
	}

	filename := result[51:]
	
	err = aws_driver.DownloadFileFromBucket(filename, "./assets/users")
	
	return controllers.NewResponse(c, http.StatusOK, "success", "user_detail updated", response.FromDomain(user_detail))
}

func (cc *UserDetailController) Delete(c echo.Context) error {
	var user_detailID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.user_detailUsecase.Delete(ctx, user_detailID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete user_detail failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "user_detail deleted", "")
}
