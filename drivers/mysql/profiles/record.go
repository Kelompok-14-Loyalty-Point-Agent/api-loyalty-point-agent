package profiles

import (
	"api-loyalty-point-agent/businesses/profiles"
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID                 uint           `json:"id" gorm:"primaryKey"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Address            string         `json:"address"`
	Age                uint           `json:"age"`
	Gender             string         `json:"gender" gorm:"type:enum('man', 'woman', 'not-selected');default:''not-selected';not_null"`
	Phone              string         `json:"phone"`
	Point              float32        `json:"point"`
	Member             string         `json:"member" gorm:"type:enum('bronze', 'silver', 'gold', 'platinum');default:'bronze';not_null"`
	TransactionMade    uint           `json:"count_transaction"`
	MonthlyTransaction uint           `json:"monthly_transaction"`
	TotalRedeem        uint           `json:"total_redeem"`
	URL                string         `json:"url"`
}

func (record *Profile) ToDomain() profiles.Domain {
	return profiles.Domain{
		ID:                 record.ID,
		CreatedAt:          record.CreatedAt,
		UpdatedAt:          record.UpdatedAt,
		DeletedAt:          record.DeletedAt,
		Address:            record.Address,
		Age:                record.Age,
		Gender:             record.Gender,
		Phone:              record.Phone,
		Point:              record.Point,
		Member:             record.Member,
		TransactionMade:    record.TransactionMade,
		TotalRedeem:        record.TotalRedeem,
		MonthlyTransaction: record.MonthlyTransaction,
		URL:                record.URL,
	}
}

func FromDomain(domain *profiles.Domain) *Profile {
	return &Profile{
		ID:                 domain.ID,
		CreatedAt:          domain.CreatedAt,
		UpdatedAt:          domain.UpdatedAt,
		DeletedAt:          domain.DeletedAt,
		Address:            domain.Address,
		Age:                domain.Age,
		Gender:             domain.Gender,
		Phone:              domain.Phone,
		Point:              domain.Point,
		Member:             domain.Member,
		TransactionMade:    domain.TransactionMade,
		TotalRedeem:        domain.TotalRedeem,
		MonthlyTransaction: domain.MonthlyTransaction,
		URL:                domain.URL,
	}
}
