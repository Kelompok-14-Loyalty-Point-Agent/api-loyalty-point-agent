package user_details

import (
	"api-loyalty-point-agent/businesses/user_details"
	"time"

	"gorm.io/gorm"
)

type UserDetail struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Member           string         `json:"member" gorm:"type:enum('bronze', 'silver', 'gold', 'platinum');default:'bronze';not_null"`
	Age              string         `json:"age"`
	Gender           string         `json:"gender" gorm:"type:enum('man', 'woman', 'not selected');default:'not selected'"`
	Address          string         `json:"address"`
	// Email            string         `json:"email" gorm:"uniqueIndex"`
	PhoneNumber      string         `json:"phone_number" gorm:"unique"`
	URL              string         `json:"url"`
	TPoint           int64          `json:"tPoint" gorm:"not_null"`
	CountTransaction int64          `json:"count_transaction" gorm:"not_null"`
	CountRedeem      int64          `json:"count_redeem" gorm:"not_null"`
}

func (rec *UserDetail) ToDomain() user_details.Domain {
	return user_details.Domain{
		ID:               rec.ID,
		CreatedAt:        rec.CreatedAt,
		UpdatedAt:        rec.UpdatedAt,
		DeletedAt:        rec.DeletedAt,
		Age:              rec.Age,
		Gender:           rec.Gender,
		Address:          rec.Address,
		PhoneNumber:      rec.PhoneNumber,
		Member:           rec.Member,
		URL:              rec.URL,
		TPoint:           rec.TPoint,
		CountTransaction: rec.CountTransaction,
		CountRedeem:      rec.CountRedeem,
	}
}

func FromDomain(domain *user_details.Domain) *UserDetail {
	return &UserDetail{
		ID:               domain.ID,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
		DeletedAt:        domain.DeletedAt,
		Age:              domain.Age,
		Gender:           domain.Gender,
		Address:          domain.Address,
		PhoneNumber:      domain.PhoneNumber,
		Member:           domain.Member,
		URL:              domain.URL,
		TPoint:           domain.TPoint,
		CountTransaction: domain.CountTransaction,
		CountRedeem:      domain.CountRedeem,
	}
}
