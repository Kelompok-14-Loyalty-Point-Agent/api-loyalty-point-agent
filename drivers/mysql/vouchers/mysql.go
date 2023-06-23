package vouchers

import (
	"api-loyalty-point-agent/businesses/redeems"
	"api-loyalty-point-agent/businesses/vouchers"
	"api-loyalty-point-agent/drivers/mysql/profiles"
	_dbRedeem "api-loyalty-point-agent/drivers/mysql/redeems"
	"context"
	"errors"
	"time"

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

	if err := cr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	vouchers := []vouchers.Domain{}

	for _, voucher := range records {
		vouchers = append(vouchers, voucher.ToDomain())
	}

	return vouchers, nil
}

func (ur *voucherRepository) GetByID(ctx context.Context, id string) (vouchers.Domain, error) {
	var voucher Voucher

	if err := ur.conn.WithContext(ctx).First(&voucher, "id = ?", id).Error; err != nil {
		return vouchers.Domain{}, err
	}

	return voucher.ToDomain(), nil

}

func (cr *voucherRepository) RedeemVoucher(ctx context.Context, redeemsDomain *redeems.Domain) (redeems.Domain, error) {
	record := _dbRedeem.FromDomain(redeemsDomain)

	var voucher Voucher

	if err := cr.conn.WithContext(ctx).First(&voucher, `id = ?`, record.VoucherID).Error; err != nil {
		return redeems.Domain{}, err
	}

	// cost
	record.Cost = voucher.Cost

	var profile profiles.Profile

	if err := cr.conn.WithContext(ctx).First(&profile, `id = ?`, record.UserID).Error; err != nil {
		return redeems.Domain{}, err
	}

	if profile.Point - voucher.Cost < 0 {
		return redeems.Domain{}, errors.New("insufficient point")
	}

	// point decrement
	profile.Point -= voucher.Cost

	// total redeem
	profile.TotalRedeem += 1

	// exchange reedem
	record.DateExchange = time.Now()

	// product
	record.Product = voucher.Product + " " + voucher.Benefit

	// payment method
	record.PaymentMethod = "tPoint"

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return redeems.Domain{}, err
	}

	err := cr.conn.WithContext(ctx).Last(&record).Error

	if err != nil {
		return redeems.Domain{}, err
	}

	if err := cr.conn.WithContext(ctx).Save(&profile).Error; err != nil {
		return redeems.Domain{}, err
	}

	return record.ToDomain(), nil
}


