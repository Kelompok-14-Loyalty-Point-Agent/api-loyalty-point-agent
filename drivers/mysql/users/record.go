package users

import (
	"api-loyalty-point-agent/businesses/users"
	"api-loyalty-point-agent/drivers/mysql/profiles"
	"api-loyalty-point-agent/drivers/mysql/redeems"
	"api-loyalty-point-agent/drivers/mysql/stock_transactions"
	"api-loyalty-point-agent/drivers/mysql/transactions"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               uint                                  `json:"id" gorm:"primaryKey"`
	CreatedAt        time.Time                             `json:"created_at"`
	UpdatedAt        time.Time                             `json:"updated_at"`
	DeletedAt        gorm.DeletedAt                        `json:"deleted_at" gorm:"index"`
	Name             string                                `json:"name"`
	Email            string                                `json:"email" gorm:"unique"`
	Password         string                                `json:"password"`
	Role             string                                `json:"role" gorm:"type:enum('admin', 'customer');default:'customer';not_null"`
	StockTransaction []stock_transactions.StockTransaction `json:"-" gorm:"foreignKey:UserID"`
	Transaction      []transactions.Transaction            `json:"-" gorm:"foreignKey:UserID"`
	Redeem           []redeems.Redeem                      `json:"-" gorm:"foreignKey:UserID"`
	ProfileID        uint                                  `json:"-" gorm:"uniqueIndex"`
	Profile          profiles.Profile                      `json:"-" gorm:"foreignKey:ProfileID"`
}

func (rec *User) ToDomain() users.Domain {
	return users.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Role:      rec.Role,
		ProfileID: rec.ProfileID,
		Profile:   rec.Profile.ToDomain(),
	}
}

func FromDomain(domain *users.Domain) *User {
	return &User{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		Role:      domain.Role,
		ProfileID: domain.ProfileID,
		Profile:   *profiles.FromDomain(&domain.Profile),
	}
}
