package providers

import (
	"api-loyalty-point-agent/businesses/providers"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/providers/request"
	"api-loyalty-point-agent/controllers/providers/response"
	"net/http"

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

func (cc *ProviderController) Create(c echo.Context) error {
	input := request.Provider{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	provider, err := cc.providerUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a provider", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "provider created", response.FromDomain(provider))
}

func (cc *ProviderController) Update(c echo.Context) error {
	var providerID string = c.Param("id")
	ctx := c.Request().Context()

	input := request.Provider{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	provider, err := cc.providerUsecase.Update(ctx, input.ToDomain(), providerID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update provider failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "provider updated", response.FromDomain(provider))
}

func (cc *ProviderController) Delete(c echo.Context) error {
	var providerID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.providerUsecase.Delete(ctx, providerID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete provider failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "provider deleted", "")
}