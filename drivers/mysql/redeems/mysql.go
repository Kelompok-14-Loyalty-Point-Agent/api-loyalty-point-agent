package redeems

import (
	redeems "api-loyalty-point-agent/businesses/redeems"
	"context"

	"gorm.io/gorm"
)

type redeemRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) redeems.Repository {
	return &redeemRepository{
		conn: conn,
	}
}

func (cr *redeemRepository) GetAll(ctx context.Context) ([]redeems.Domain, error) {
	var records []Redeem

	// if err := cr.conn.WithContext(ctx).Preload("Profile").Find(&records).Error; err != nil {
	// 	return nil, err
	// }

	if err := cr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	vouchers := []redeems.Domain{}

	for _, voucher := range records {
		vouchers = append(vouchers, voucher.ToDomain())
	}

	return vouchers, nil
}

// =========================================================
func (cr *redeemRepository) RedeemVoucher(ctx context.Context, voucherDomain *redeems.Domain) (redeems.Domain, error) {
	record := FromDomain(voucherDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return redeems.Domain{}, err
	}

	err := cr.conn.WithContext(ctx).Last(&record).Error

	if err != nil {
		return redeems.Domain{}, err
	}

	return record.ToDomain(), nil
}
