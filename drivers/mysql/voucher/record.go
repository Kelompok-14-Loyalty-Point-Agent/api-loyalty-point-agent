package voucher

import (
	"api-loyalty-point-agent/businesses/voucher"
	// "api-loyalty-point-agent/drivers/mysql/stock_details"
	// "api-loyalty-point-agent/drivers/mysql/stocks"
	// "api-loyalty-point-agent/drivers/mysql/stocks"
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

func (rec *Voucher) ToDomain() voucher.Domain {
	return voucher.Domain{
		Title: rec.Title,
		Url:   rec.Url,
		Cost:  rec.Cost,
		Point: rec.Point,
		// Status:         domain.Status,
		// Date_Exhange: req.Date_Exhange,
	}
}

func FromDomain(domain *voucher.Domain) *Voucher {
	return &Voucher{
		ID:    domain.ID,
		Title: domain.Title,
		Url:   domain.Url,
		Cost:  domain.Cost,
		Point: domain.Point,
		// Status:         domain.Status,
		// Date_Exhange: domain.Date_Exhange,
	}
}
