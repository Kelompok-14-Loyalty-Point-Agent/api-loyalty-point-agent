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

	transactionDomain "api-loyalty-point-agent/businesses/transactions"
	transactionDB "api-loyalty-point-agent/drivers/mysql/transactions"

	stock_transactionDomain "api-loyalty-point-agent/businesses/stock_transactions"
	stock_transactionDB "api-loyalty-point-agent/drivers/mysql/stock_transactions"

	profileDomain "api-loyalty-point-agent/businesses/profiles"
	profileDB "api-loyalty-point-agent/drivers/mysql/profiles"

	voucherDomain "api-loyalty-point-agent/businesses/vouchers"
	voucherDB "api-loyalty-point-agent/drivers/mysql/vouchers"

	redeemsDomain "api-loyalty-point-agent/businesses/redeems"
	redeemsDB "api-loyalty-point-agent/drivers/mysql/redeems"

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

func NewTransactionRepository(conn *gorm.DB) transactionDomain.Repository {
	return transactionDB.NewMySQLRepository(conn)
}

func NewStockTransactionRepository(conn *gorm.DB) stock_transactionDomain.Repository {
	return stock_transactionDB.NewMySQLRepository(conn)
}

func NewProfileRepository(conn *gorm.DB) profileDomain.Repository {
	return profileDB.NewMySQLRepository(conn)
}

func NewVoucherRepository(conn *gorm.DB) voucherDomain.Repository {
	return voucherDB.NewMySQLRepository(conn)
}

func NewRedeemRepository(conn *gorm.DB) redeemsDomain.Repository {
	return redeemsDB.NewMySQLRepository(conn)
}
