package response

import (
	voucher "api-loyalty-point-agent/businesses/vouchers"
)

type Voucher struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Product string `json:"product"`
	Benefit string `json:"benefit"`
	Cost    uint   `json:"cost"`
}

func FromDomain(domain voucher.Domain) Voucher {
	return Voucher{
		ID:      domain.ID,
		Product: domain.Product,
		Benefit: domain.Benefit,
		Cost:    domain.Cost,
	}
}
