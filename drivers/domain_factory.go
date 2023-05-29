package drivers

import (

	userDomain "api-loyalty-point-agent/businesses/users"
	userDB "api-loyalty-point-agent/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}