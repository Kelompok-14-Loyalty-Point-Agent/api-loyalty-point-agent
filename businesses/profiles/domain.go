package profiles

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID                 uint
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
	Name               string
	Address            string
	Age                uint
	Gender             string
	Phone              string
	Point              float64
	Member             string
	TransactionMade    uint
	MonthlyTransaction uint
	TotalRedeem        uint
	URL                string
}
