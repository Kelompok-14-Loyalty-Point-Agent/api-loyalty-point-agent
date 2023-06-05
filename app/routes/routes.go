package routes

import (
	"api-loyalty-point-agent/app/middlewares"

	providers "api-loyalty-point-agent/controllers/providers"
	stock_details "api-loyalty-point-agent/controllers/stock_details"
	stocks "api-loyalty-point-agent/controllers/stocks"
	users "api-loyalty-point-agent/controllers/users"
	user_details "api-loyalty-point-agent/controllers/user_details"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware      echo.MiddlewareFunc
	JWTMiddleware         echojwt.Config
	AuthController        users.AuthController
	ProviderController    providers.ProviderController
	StockController       stocks.StockController
	StockDetailController stock_details.StockDetailController
	UserDetailController user_details.UserDetailController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)
	auth := e.Group("auth")
	auth.POST("/register", cl.AuthController.Register)
	auth.POST("/login", cl.AuthController.Login)

	users := e.Group("/users", echojwt.WithConfig(cl.JWTMiddleware))
	users.Use(middlewares.VerifyToken)
	users.POST("/logout", cl.AuthController.Logout)
	users.GET("/customers", cl.AuthController.GetAllCustomers)

	user_details := e.Group("/users/details", echojwt.WithConfig(cl.JWTMiddleware))
	user_details.Use(middlewares.VerifyToken)
	user_details.GET("", cl.UserDetailController.GetAll)
	user_details.GET("", cl.UserDetailController.GetByID)
	user_details.PUT("/:id", cl.UserDetailController.Update)
	user_details.DELETE("/:id", cl.UserDetailController.Delete)

	providers := e.Group("/providers", echojwt.WithConfig(cl.JWTMiddleware))
	providers.Use(middlewares.VerifyToken)
	providers.GET("", cl.ProviderController.GetAll)
	// read image from bucket
	providers.GET("/read", cl.ProviderController.ReadFile)
	// download image from bucket
	providers.GET("/image/download", cl.ProviderController.DownloadFile)
	providers.GET("/:id", cl.ProviderController.GetByID)
	providers.POST("", cl.ProviderController.Create)
	providers.PUT("/:id", cl.ProviderController.Update)
	providers.DELETE("/:id", cl.ProviderController.Delete)

	stocks := e.Group("/stocks", echojwt.WithConfig(cl.JWTMiddleware))
	stocks.Use(middlewares.VerifyToken)
	stocks.GET("", cl.StockController.GetAll)
	stocks.GET("/:id", cl.StockController.GetByID)
	stocks.POST("", cl.StockController.Create)
	stocks.PUT("/:id", cl.StockController.Update)
	stocks.DELETE("/:id", cl.StockController.Delete)

	stock_details := e.Group("/stocks/details", echojwt.WithConfig(cl.JWTMiddleware))
	stock_details.Use(middlewares.VerifyToken)
	stock_details.GET("", cl.StockDetailController.GetAll)
	stock_details.GET("/:id", cl.StockDetailController.GetByID)
	stock_details.POST("", cl.StockDetailController.Create)
	stock_details.PUT("/:id", cl.StockDetailController.Update)
	stock_details.DELETE("/:id", cl.StockDetailController.Delete)

	// admin := e.Group("/admin", echojwt.WithConfig(cl.JWTMiddleware))
	// admin.Use(middlewares.VerifyToken)
	// users.GET("/stock", cl.AuthController.GetAllCustomers)

}
