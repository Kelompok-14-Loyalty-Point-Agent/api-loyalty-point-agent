package drivers

import (
	userDomain "api-loyalty-point-agent/businesses/users"
	userDB "api-loyalty-point-agent/drivers/mysql/users"

	providerDomain "api-loyalty-point-agent/businesses/providers"
	providerDB "api-loyalty-point-agent/drivers/mysql/providers"

	stock_detailDomain "api-loyalty-point-agent/businesses/stock_details"
	stock_detailDB "api-loyalty-point-agent/drivers/mysql/stock_details"

	stockDomain "api-loyalty-point-agent/businesses/stocks"
	stockDB "api-loyalty-point-agent/drivers/mysql/stocks"

	stock_transactionDomain "api-loyalty-point-agent/businesses/stock_transactions"
	stock_transactionDB "api-loyalty-point-agent/drivers/mysql/stock_transactions"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}

func NewProviderRepository(conn *gorm.DB) providerDomain.Repository {
	return providerDB.NewMySQLRepository(conn)
}

func NewStockDetailRepository(conn *gorm.DB) stock_detailDomain.Repository {
	return stock_detailDB.NewMySQLRepository(conn)
}

func NewStockRepository(conn *gorm.DB) stockDomain.Repository {
	return stockDB.NewMySQLRepository(conn)
}

func NewStockTransactionRepository(conn *gorm.DB) stock_transactionDomain.Repository {
	return stock_transactionDB.NewMySQLRepository(conn)
}
