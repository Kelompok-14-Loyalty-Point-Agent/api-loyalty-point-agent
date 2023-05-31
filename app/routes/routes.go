package routes

import (
	"api-loyalty-point-agent/app/middlewares"
	// "api-loyalty-point-agent/businesses/providers"
	providers "api-loyalty-point-agent/controllers/providers"
	stocks "api-loyalty-point-agent/controllers/stocks"

	users "api-loyalty-point-agent/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware   echo.MiddlewareFunc
	JWTMiddleware      echojwt.Config
	AuthController     users.AuthController
	ProviderController providers.ProviderController
	StockController    stocks.StockController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)
	user := e.Group("auth")
	user.POST("/register", cl.AuthController.Register)
	user.POST("/login", cl.AuthController.Login)

	users := e.Group("/users", echojwt.WithConfig(cl.JWTMiddleware))
	users.Use(middlewares.VerifyToken)
	users.POST("/logout", cl.AuthController.Logout)
	users.GET("/customers", cl.AuthController.GetAllCustomers)

	providers := e.Group("/provider", echojwt.WithConfig(cl.JWTMiddleware))
	providers.Use(middlewares.VerifyToken)
	providers.GET("/providers", cl.ProviderController.GetAllProvider)
	providers.GET("/providers/:id", cl.ProviderController.GetByIDProvider)
	providers.POST("/providers", cl.ProviderController.CreateProvider)
	providers.PUT("/providers/:id", cl.ProviderController.UpdateProvider)
	providers.DELETE("/providers/:id", cl.ProviderController.DeleteProvider)

	stocks := e.Group("/stock", echojwt.WithConfig(cl.JWTMiddleware))
	stocks.Use(middlewares.VerifyToken)
	stocks.GET("/stocks", cl.StockController.GetAll)
	stocks.GET("/stocks/:id", cl.StockController.GetByID)
	stocks.POST("/stocks", cl.StockController.Create)
	stocks.PUT("/stocks/:id", cl.StockController.Update)
	stocks.DELETE("/stocks/:id", cl.StockController.Delete)

	// admin := e.Group("/admin", echojwt.WithConfig(cl.JWTMiddleware))
	// admin.Use(middlewares.VerifyToken)
	// users.GET("/stock", cl.AuthController.GetAllCustomers)

}
