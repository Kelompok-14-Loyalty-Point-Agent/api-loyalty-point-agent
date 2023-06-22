package response

import (
	"api-loyalty-point-agent/businesses/voucher"
)

type Voucher struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Cost  uint   `json:"cost"`
	Point uint   `json:"point"`
	// Expired	uint
	// Date_Exhange time.Time `json:"date_exhange"`
}

func FromDomain(domain voucher.Domain) Voucher {
	return Voucher{
		ID:    domain.ID,
		Title: domain.Title,
		Url:   domain.Url,
		Cost:  domain.Cost,
		Point: domain.Point,
		// Status:         domain.Status,
		// Date_Exhange: domain.Date_Exhange,
	}
}
