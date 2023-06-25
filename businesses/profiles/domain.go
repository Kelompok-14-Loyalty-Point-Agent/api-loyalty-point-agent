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
	Address            string
	Age                uint
	Gender             string
	Phone              string
	Point              float32
	Member             string
	TransactionMade    uint
	MonthlyTransaction uint
	TotalRedeem        uint
	URL                string
}
