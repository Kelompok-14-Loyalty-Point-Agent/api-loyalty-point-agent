package vouchers

import (
	"api-loyalty-point-agent/businesses/vouchers"
	"context"

	"gorm.io/gorm"
)

type voucherRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) vouchers.Repository {
	return &voucherRepository{
		conn: conn,
	}
}

func (cr *voucherRepository) GetAll(ctx context.Context) ([]vouchers.Domain, error) {
	var records []Voucher

	// if err := cr.conn.WithContext(ctx).Preload("Profile").Find(&records).Error; err != nil {
	// 	return nil, err
	// }

	if err := cr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	vouchers := []vouchers.Domain{}

	for _, voucher := range records {
		vouchers = append(vouchers, voucher.ToDomain())
	}

	return vouchers, nil
}

// =========================================================
func (cr *voucherRepository) Create(ctx context.Context, voucherDomain *vouchers.Domain) (vouchers.Domain, error) {
	record := FromDomain(voucherDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return vouchers.Domain{}, err
	}

	err := cr.conn.WithContext(ctx).Last(&record).Error

	if err != nil {
		return vouchers.Domain{}, err
	}

	return record.ToDomain(), nil
}
