package providers

import (
	"api-loyalty-point-agent/businesses/providers"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/providers/response"
	"encoding/json"

	// "io"

	aws_driver "api-loyalty-point-agent/drivers/aws"
	// "api-loyalty-point-agent/utils"
	"net/http"
	// "os"
	// "path/filepath"

	"github.com/labstack/echo/v4"
)

type ProviderController struct {
	providerUsecase providers.Usecase
}

func NewProviderController(providerUC providers.Usecase) *ProviderController {
	return &ProviderController{
		providerUsecase: providerUC,
	}
}

func (cc *ProviderController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	providersData, err := cc.providerUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	providers := []response.Provider{}

	for _, provider := range providersData {
		providers = append(providers, response.FromDomain(provider))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all providers", providers)
}

func (cc *ProviderController) GetByID(c echo.Context) error {
	var providerID string = c.Param("id")
	ctx := c.Request().Context()

	provider, err := cc.providerUsecase.GetByID(ctx, providerID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "provider not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "provider found", response.FromDomain(provider))
}

func (cc *ProviderController) ReadFile(c echo.Context) error {
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

func (cc *ProviderController) DownloadFile(c echo.Context) error {
	name := c.FormValue("filename")

	err := aws_driver.DownloadFileFromBucket(name, "./assets/providers/")

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "downloaded data from bucket", "")
}

// func (cc *ProviderController) Create(c echo.Context) error {
// 	name := c.FormValue("name")
// 	file, err := c.FormFile("picture")

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "failed to upload file", "")
// 	}

// 	src, err := file.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer src.Close()

// 	result, err := aws_driver.UploadFileToBucket(file.Filename, src)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
// 	}

// 	input := request.Provider{
// 		Name: name,
// 		URL: result,
// 	}

// 	jsonBody, err := json.Marshal(&input)

// 	if err != nil {
// 		return err
// 	}

// 	err = json.Unmarshal(jsonBody, &input)
// 	if err != nil {
// 		return err
// 	}

// 	ctx := c.Request().Context()

// 	provider, err := cc.providerUsecase.Create(ctx, input.ToDomain())
// 	if err != nil {
// 		return err
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, "success", "file uploaded successfully", response.FromDomain(provider))
// }

