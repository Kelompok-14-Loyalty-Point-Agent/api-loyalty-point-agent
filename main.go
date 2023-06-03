package main

import (
	_driverFactory "api-loyalty-point-agent/drivers"
	"api-loyalty-point-agent/utils"

	_userUseCase "api-loyalty-point-agent/businesses/users"
	_userController "api-loyalty-point-agent/controllers/users"

	_providerUseCase "api-loyalty-point-agent/businesses/providers"
	_providerController "api-loyalty-point-agent/controllers/providers"

	_stockUseCase "api-loyalty-point-agent/businesses/stocks"
	_stockController "api-loyalty-point-agent/controllers/stocks"

	_stock_detailUseCase "api-loyalty-point-agent/businesses/stock_details"
	_stock_detailController "api-loyalty-point-agent/controllers/stock_details"

	_dbDriver "api-loyalty-point-agent/drivers/mysql"

	_middleware "api-loyalty-point-agent/app/middlewares"
	_routes "api-loyalty-point-agent/app/routes"

	_ "api-loyalty-point-agent/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title API Loyalty Point Agent
// @version 1.0
// @description Berikut API Loyalty Point Agent.
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /

func main() {

	configDB := _dbDriver.DBConfig{
		DB_USERNAME: utils.GetConfig("DB_USERNAME"),
		DB_PASSWORD: utils.GetConfig("DB_PASSWORD"),
		DB_HOST:     utils.GetConfig("DB_HOST"),
		DB_PORT:     utils.GetConfig("DB_PORT"),
		DB_NAME:     utils.GetConfig("DB_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.MigrateDB(db)

	_dbDriver.SeedAdmin(db)

	configJWT := _middleware.JWTConfig{
		SecretKey:       utils.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	configLogger := _middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	e := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUseCase.NewUserUseCase(userRepo, &configJWT)
	userCtrl := _userController.NewAuthController(userUsecase)

	providerRepo := _driverFactory.NewProviderRepository(db)
	providerUsecase := _providerUseCase.NewProviderUseCase(providerRepo, &configJWT)
	providerCtrl := _providerController.NewProviderController(providerUsecase)

	stockRepo := _driverFactory.NewStockRepository(db)
	stockUsecase := _stockUseCase.NewStockUseCase(stockRepo, &configJWT)
	stockCtrl := _stockController.NewStockController(stockUsecase)

	stock_detailRepo := _driverFactory.NewStockDetailRepository(db)
	stock_detailUsecase := _stock_detailUseCase.NewStockDetailUseCase(stock_detailRepo, &configJWT)
	stock_detailCtrl := _stock_detailController.NewStockDetailController(stock_detailUsecase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware:      configLogger.Init(),
		JWTMiddleware:         configJWT.Init(),
		AuthController:        *userCtrl,
		StockController:       *stockCtrl,
		ProviderController:    *providerCtrl,
		StockDetailController: *stock_detailCtrl,
	}

	handleSwagger := echoSwagger.WrapHandler
	e.GET("/swagger/*", handleSwagger)

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "API Is Active")
	// })

	routesInit.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
