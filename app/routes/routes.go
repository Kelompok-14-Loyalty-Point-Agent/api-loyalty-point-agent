package routes

import (
	"api-loyalty-point-agent/app/middlewares"

	profiles "api-loyalty-point-agent/controllers/profiles"
	providers "api-loyalty-point-agent/controllers/providers"
	stock_details "api-loyalty-point-agent/controllers/stock_details"
	stock_transactions "api-loyalty-point-agent/controllers/stock_transactions"
	stocks "api-loyalty-point-agent/controllers/stocks"
	transactions "api-loyalty-point-agent/controllers/transactions"
	users "api-loyalty-point-agent/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware           echo.MiddlewareFunc
	JWTMiddleware              echojwt.Config
	AuthController             users.AuthController
	ProviderController         providers.ProviderController
	StockController            stocks.StockController
	StockDetailController      stock_details.StockDetailController
	StockTransactionController stock_transactions.StockTransactionController
	TransactionController      transactions.TransactionController
	ProfileController          profiles.ProfileController
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

	users.GET("/:id", cl.AuthController.GetByID)
	users.PUT("/profiles/customer/:id", cl.AuthController.UpdateProfileCustomer)
	users.PUT("/profiles/admin/:id", cl.AuthController.UpdateProfileAdmin)
	users.PUT("/profiles/password/:id", cl.AuthController.ChangePassword)
	users.DELETE("/customers/:id", cl.AuthController.DeleteCustomer)

	transactions := e.Group("/transactions", echojwt.WithConfig(cl.JWTMiddleware))
	transactions.Use(middlewares.VerifyToken)
	transactions.GET("", cl.TransactionController.GetAll)
	transactions.GET("/:id", cl.TransactionController.GetByID)
	transactions.POST("", cl.TransactionController.Create)
	transactions.GET("/users/:id", cl.TransactionController.GetAllByUserID)
	transactions.PUT("/:id", cl.TransactionController.UpdatePoint)

	providers := e.Group("/providers", echojwt.WithConfig(cl.JWTMiddleware))
	providers.Use(middlewares.VerifyToken)
	providers.GET("", cl.ProviderController.GetAll)
	// read image from bucket
	providers.GET("/read", cl.ProviderController.ReadFile)
	// download image from bucket
	providers.GET("/image/download", cl.ProviderController.DownloadFile)
	providers.GET("/:id", cl.ProviderController.GetByID)

	stocks := e.Group("/stocks", echojwt.WithConfig(cl.JWTMiddleware))
	stocks.Use(middlewares.VerifyToken)
	stocks.GET("", cl.StockController.GetAll)
	stocks.GET("/:id", cl.StockController.GetByID)
	// add stock and create stock transaction data
	stocks.POST("/add", cl.StockController.AddStock)

	stock_detail := e.Group("/stocks/details", echojwt.WithConfig(cl.JWTMiddleware))
	stock_detail.Use(middlewares.VerifyToken)
	stock_detail.GET("", cl.StockDetailController.GetAll)
	stock_detail.GET("/:id", cl.StockDetailController.GetByID)
	stock_detail.POST("", cl.StockDetailController.Create)
	stock_detail.PUT("/:id", cl.StockDetailController.Update)
	stock_detail.DELETE("/:id", cl.StockDetailController.Delete)
	stock_detail.GET("/bystocks/:id", cl.StockDetailController.GetAllByStockID)

	stock_transaction := e.Group("/stocks/transactions", echojwt.WithConfig(cl.JWTMiddleware))
	stock_transaction.Use(middlewares.VerifyToken)
	stock_transaction.GET("", cl.StockTransactionController.GetAll)
	stock_transaction.GET("/:id", cl.StockTransactionController.GetByID)

	// admin := e.Group("/admin", echojwt.WithConfig(cl.JWTMiddleware))
	// admin.Use(middlewares.VerifyToken)
	// users.GET("/stock", cl.AuthController.GetAllCustomers)

	profiles := e.Group("/profiles", echojwt.WithConfig(cl.JWTMiddleware))
	profiles.Use(middlewares.VerifyToken)
	profiles.GET("", cl.ProfileController.GetAll)
	profiles.GET("/:id", cl.ProfileController.GetByID)
	profiles.PUT("/:id", cl.ProfileController.Update)
	profiles.DELETE("/:id", cl.ProfileController.Delete)
}
