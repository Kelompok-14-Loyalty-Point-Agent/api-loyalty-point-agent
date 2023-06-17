package stock_details

import (
	"api-loyalty-point-agent/businesses/stock_details"
	"context"

	"gorm.io/gorm"
)

type stock_detailRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) stock_details.Repository {
	return &stock_detailRepository{
		conn: conn,
	}
}

func (cr *stock_detailRepository) GetAll(ctx context.Context) ([]stock_details.Domain, error) {
	var records []StockDetail

	if err := cr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	stock_details := []stock_details.Domain{}

	for _, stock_detail := range records {
		stock_details = append(stock_details, stock_detail.ToDomain())
	}

	return stock_details, nil
}

func (cr *stock_detailRepository) GetByID(ctx context.Context, id string) (stock_details.Domain, error) {
	var stock_detail StockDetail

	if err := cr.conn.WithContext(ctx).First(&stock_detail, "id = ?", id).Error; err != nil {
		return stock_details.Domain{}, err
	}

	return stock_detail.ToDomain(), nil
}

func (cr *stock_detailRepository) Create(ctx context.Context, stock_detailDomain *stock_details.Domain) (stock_details.Domain, error) {
	record := FromDomain(stock_detailDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return stock_details.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return stock_details.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (cr *stock_detailRepository) Update(ctx context.Context, stock_detailDomain *stock_details.Domain, id string) (stock_details.Domain, error) {
	stock_detail, err := cr.GetByID(ctx, id)

	if err != nil {
		return stock_details.Domain{}, err
	}

	updatedStockDetail := FromDomain(&stock_detail)

	if updatedStockDetail.Price != stock_detailDomain.Price {
		updatedStockDetail.Price = stock_detailDomain.Price
	}

	if updatedStockDetail.Quantity != stock_detailDomain.Quantity {
		updatedStockDetail.Quantity = stock_detailDomain.Quantity
	}
	
	if updatedStockDetail.Stock != stock_detailDomain.Stock {
		updatedStockDetail.Stock = stock_detailDomain.Stock
	}
	
	if err := cr.conn.WithContext(ctx).Save(&updatedStockDetail).Error; err != nil {
		return stock_details.Domain{}, err
	}

	return updatedStockDetail.ToDomain(), nil
}

func (cr *stock_detailRepository) Delete(ctx context.Context, id string) error {
	stock_detail, err := cr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedStockDetail := FromDomain(&stock_detail)

	if err := cr.conn.WithContext(ctx).Unscoped().Delete(&deletedStockDetail).Error; err != nil {
		return err
	}

	return nil
}
