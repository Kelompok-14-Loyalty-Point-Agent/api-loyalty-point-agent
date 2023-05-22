package drivers

import (

	customerDomain "api-loyalty-point-agent/businesses/customers"
	customerDB "api-loyalty-point-agent/drivers/mysql/customers"

	"gorm.io/gorm"
)

func NewCustomerRepository(conn *gorm.DB) customerDomain.Repository {
	return customerDB.NewMySQLRepository(conn)
}