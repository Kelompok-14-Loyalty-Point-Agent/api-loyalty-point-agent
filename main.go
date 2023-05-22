package main

import (
	_driverFactory "api-loyalty-point-agent/drivers"
	"api-loyalty-point-agent/utils"

	_customerUseCase "api-loyalty-point-agent/businesses/customers"
	_customerController "api-loyalty-point-agent/controllers/customers"

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

	configJWT := _middleware.JWTConfig{
		SecretKey:       utils.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	configLogger := _middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	e := echo.New()

	customerRepo := _driverFactory.NewCustomerRepository(db)
	customerUsecase := _customerUseCase.NewCustomerUseCase(customerRepo, &configJWT)
	customerCtrl := _customerController.NewAuthController(customerUsecase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware:   configLogger.Init(),
		JWTMiddleware:      configJWT.Init(),
		AuthController:     *customerCtrl,
	}

	handleSwagger := echoSwagger.WrapHandler
	e.GET("/swagger/*", handleSwagger)

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "API Is Active")
	// })

	routesInit.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
