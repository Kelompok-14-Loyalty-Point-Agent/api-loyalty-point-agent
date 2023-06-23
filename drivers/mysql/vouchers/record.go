package vouchers

import (
	"api-loyalty-point-agent/businesses/vouchers"
	"api-loyalty-point-agent/drivers/mysql/redeems"
)

type Voucher struct {
	ID      uint             `json:"id" gorm:"primaryKey"`
	Product string           `json:"product"`
	Benefit string           `json:"benefit"`
	Cost    float32          `json:"cost"`
	Redeem  []redeems.Redeem `json:"-" gorm:"foreignKey:VoucherID"`
}

func (rec *Voucher) ToDomain() vouchers.Domain {
	return vouchers.Domain{
		ID:      rec.ID,
		Product: rec.Product,
		Benefit: rec.Benefit,
		Cost:    rec.Cost,
	}
}

func FromDomain(domain *vouchers.Domain) *Voucher {
	return &Voucher{
		ID:      domain.ID,
		Product: domain.Product,
		Benefit: domain.Benefit,
		Cost:    domain.Cost,
	}
}
