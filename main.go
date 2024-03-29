package main

import (
	_driverFactory "api-loyalty-point-agent/drivers"
	"api-loyalty-point-agent/utils"
	"net/http"

	_userUseCase "api-loyalty-point-agent/businesses/users"
	_userController "api-loyalty-point-agent/controllers/users"

	_stockUseCase "api-loyalty-point-agent/businesses/stocks"
	_stockController "api-loyalty-point-agent/controllers/stocks"

	_stock_detailUseCase "api-loyalty-point-agent/businesses/stock_details"
	_stock_detailController "api-loyalty-point-agent/controllers/stock_details"

	_stock_transactionUseCase "api-loyalty-point-agent/businesses/stock_transactions"
	_stock_transactionController "api-loyalty-point-agent/controllers/stock_transactions"

	_transactionUseCase "api-loyalty-point-agent/businesses/transactions"
	_transactionController "api-loyalty-point-agent/controllers/transactions"

	_voucherUseCase "api-loyalty-point-agent/businesses/vouchers"
	_voucherController "api-loyalty-point-agent/controllers/vouchers"

	_redeemsUseCase "api-loyalty-point-agent/businesses/redeems"
	_redeemsController "api-loyalty-point-agent/controllers/redeems"

	_dbDriver "api-loyalty-point-agent/drivers/mysql"

	_middleware "api-loyalty-point-agent/app/middlewares"
	_routes "api-loyalty-point-agent/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

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

	_dbDriver.SeedCustomer(db)

	_dbDriver.SeedProvider(db)

	_dbDriver.SeedVoucher(db)

	_dbDriver.SeedStockDetail(db)

	_dbDriver.SeedStock(db)

	_dbDriver.SeedTransaction(db)

	configJWT := _middleware.JWTConfig{
		SecretKey:       utils.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	configLogger := _middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
		AllowHeaders:     []string{"Authorization", "Content-Type", "Origin"},
	}))

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUseCase.NewUserUseCase(userRepo, &configJWT)
	userCtrl := _userController.NewAuthController(userUsecase)

	stockRepo := _driverFactory.NewStockRepository(db)
	stockUsecase := _stockUseCase.NewStockUseCase(stockRepo, &configJWT)
	stockCtrl := _stockController.NewStockController(stockUsecase)

	stock_detailRepo := _driverFactory.NewStockDetailRepository(db)
	stock_detailUsecase := _stock_detailUseCase.NewStockDetailUseCase(stock_detailRepo, &configJWT)
	stock_detailCtrl := _stock_detailController.NewStockDetailController(stock_detailUsecase)

	stock_transactionRepo := _driverFactory.NewStockTransactionRepository(db)
	stock_transactionUsecase := _stock_transactionUseCase.NewStockTransactionUseCase(stock_transactionRepo, &configJWT)
	stock_transactionCtrl := _stock_transactionController.NewStockTransactionController(stock_transactionUsecase)

	transactionRepo := _driverFactory.NewTransactionRepository(db)
	transactionUsecase := _transactionUseCase.NewTransactionUseCase(transactionRepo, &configJWT)
	transactionCtrl := _transactionController.NewTransactionController(transactionUsecase)

	voucherRepo := _driverFactory.NewVoucherRepository(db)
	voucherUsecase := _voucherUseCase.NewVoucherUseCase(voucherRepo, &configJWT)
	voucherCtrl := _voucherController.NewVoucherController(voucherUsecase)

	redeemsRepo := _driverFactory.NewRedeemRepository(db)
	redeemsUsecase := _redeemsUseCase.NewRedeemUseCase(redeemsRepo, &configJWT)
	redeemsCtrl := _redeemsController.NewRedeemController(redeemsUsecase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware:           configLogger.Init(),
		JWTMiddleware:              configJWT.Init(),
		AuthController:             *userCtrl,
		StockController:            *stockCtrl,
		StockDetailController:      *stock_detailCtrl,
		StockTransactionController: *stock_transactionCtrl,
		TransactionController:      *transactionCtrl,
		VoucherController:          *voucherCtrl,
		RedeemsController:          *redeemsCtrl,
	}

	handleSwagger := echoSwagger.WrapHandler
	e.GET("/swagger/*", handleSwagger)

	routesInit.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
