package stocks

import (
	"api-loyalty-point-agent/businesses/stocks"
	"context"

	"gorm.io/gorm"
)

type stockRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) stocks.Repository {
	return &stockRepository{
		conn: conn,
	}
}

func (cr *stockRepository) GetAll(ctx context.Context) ([]stocks.Domain, error) {
	var records []Stock

	if err := cr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	stocks := []stocks.Domain{}

	for _, stock := range records {
		stocks = append(stocks, stock.ToDomain())
	}

	return stocks, nil
}

func (cr *stockRepository) GetByID(ctx context.Context, id string) (stocks.Domain, error) {
	var stock Stock

	if err := cr.conn.WithContext(ctx).First(&stock, "id = ?", id).Error; err != nil {
		return stocks.Domain{}, err
	}

	return stock.ToDomain(), nil
}

func (cr *stockRepository) Create(ctx context.Context, stockDomain *stocks.Domain) (stocks.Domain, error) {
	record := FromDomain(stockDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return stocks.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return stocks.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (cr *stockRepository) Update(ctx context.Context, stockDomain *stocks.Domain, id string) (stocks.Domain, error) {
	stock, err := cr.GetByID(ctx, id)

	if err != nil {
		return stocks.Domain{}, err
	}

	updatedStock := FromDomain(&stock)

	if updatedStock.Type != stockDomain.Type{
		updatedStock.Type = stockDomain.Type
	}

	if updatedStock.TotalStock != stockDomain.TotalStock {
		updatedStock.TotalStock = stockDomain.TotalStock
	}
	
	if err := cr.conn.WithContext(ctx).Save(&updatedStock).Error; err != nil {
		return stocks.Domain{}, err
	}

	return updatedStock.ToDomain(), nil
}

func (cr *stockRepository) Delete(ctx context.Context, id string) error {
	stock, err := cr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedStock := FromDomain(&stock)

	if err := cr.conn.WithContext(ctx).Unscoped().Delete(&deletedStock).Error; err != nil {
		return err
	}

	return nil
}

