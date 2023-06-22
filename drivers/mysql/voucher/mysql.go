package voucher

import (
	"api-loyalty-point-agent/businesses/voucher"
	"context"

	"gorm.io/gorm"
)

type voucherRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) voucher.Repository {
	return &voucherRepository{
		conn: conn,
	}
}

func (cr *voucherRepository) GetAll(ctx context.Context) ([]voucher.Domain, error) {
	var records []Voucher

	// if err := cr.conn.WithContext(ctx).Preload("StockDetails").Find(&records).Error; err != nil {
	// 	return nil, err
	// }
	if err := cr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	vouchers := []voucher.Domain{}

	for _, voucher := range records {
		vouchers = append(vouchers, voucher.ToDomain())
	}

	return vouchers, nil
}

func (cr *voucherRepository) RedeemVoucher(ctx context.Context, voucherDomain *voucher.Domain) (voucher.Domain, error) {
	record := FromDomain(voucherDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return voucher.Domain{}, err
	}

	err := cr.conn.WithContext(ctx).Last(&record).Error

	if err != nil {
		return voucher.Domain{}, err
	}

	return record.ToDomain(), nil
}
