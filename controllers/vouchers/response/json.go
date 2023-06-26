package response

import (
	"api-loyalty-point-agent/businesses/vouchers"
)

type Voucher struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	Product string  `json:"product"`
	Benefit string  `json:"benefit"`
	Cost    float32 `json:"cost"`
}

func FromDomain(domain vouchers.Domain) Voucher {
	return Voucher{
		ID:      domain.ID,
		Product: domain.Product,
		Benefit: domain.Benefit,
		Cost:    domain.Cost,
	}
}
