package response

import (
	"api-loyalty-point-agent/businesses/user_details"
	"time"

	"gorm.io/gorm"
)

type UserDetail struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
	Member           string         `json:"member"`
	TPoint           int64          `json:"tPoint"`
	Age              string         `json:"age" `
	Gender           string         `json:"gender"`
	Address          string         `json:"address"`
	Email            string         `json:"email"`
	PhoneNumber      string         `json:"phone_number"`
	URL              string         `json:"url"`
	CountTransaction int64          `json:"count_transaction"`
	CountRedeem      int64          `json:"count_redeem"`
}

func FromDomain(domain user_details.Domain) UserDetail {
	return UserDetail{
		ID:               domain.ID,
		Member:           domain.Member,
		TPoint:           domain.TPoint,
		Age:              domain.Age,
		Gender:           domain.Gender,
		Address:          domain.Address,
		Email:            domain.Email,
		PhoneNumber:      domain.PhoneNumber,
		URL:              domain.URL,
		CountTransaction: domain.CountTransaction,
		CountRedeem:      domain.CountRedeem,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
		DeletedAt:        domain.DeletedAt,
	}
}
