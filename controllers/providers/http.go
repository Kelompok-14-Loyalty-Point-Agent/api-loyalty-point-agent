package providers

import (
	"api-loyalty-point-agent/app/middlewares"
	"api-loyalty-point-agent/businesses/providers"
	"api-loyalty-point-agent/controllers"
	"api-loyalty-point-agent/controllers/providers/request"
	"api-loyalty-point-agent/controllers/providers/response"
	"net/http"

	"github.com/golang-jwt/jwt/v4"

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

func (cc *ProviderController) GetAllProvider(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	ctx := c.Request().Context()

	providersData, err := cc.providerUsecase.GetAllProvider(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	providers := []response.Provider{}

	for _, provider := range providersData {
		providers = append(providers, response.FromDomain(provider))
	}

	isListed := middlewares.CheckToken(token.Raw)

	if !isListed {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "invalid token", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all providers", providers)
}

func (cc *ProviderController) GetByIDProvider(c echo.Context) error {
	var providerID string = c.Param("id")
	ctx := c.Request().Context()

	provider, err := cc.providerUsecase.GetByIDProvider(ctx, providerID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "provider not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "provider found", response.FromDomain(provider))
}

func (cc *ProviderController) CreateProvider(c echo.Context) error {
	input := request.Provider{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	provider, err := cc.providerUsecase.CreateProvider(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a provider", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "provider created", response.FromDomain(provider))
}

func (cc *ProviderController) UpdateProvider(c echo.Context) error {
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

	provider, err := cc.providerUsecase.UpdateProvider(ctx, input.ToDomain(), providerID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update provider failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "provider updated", response.FromDomain(provider))
}

func (cc *ProviderController) DeleteProvider(c echo.Context) error {
	var providerID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.providerUsecase.DeleteProvider(ctx, providerID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete provider failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "provider deleted", "")
}
