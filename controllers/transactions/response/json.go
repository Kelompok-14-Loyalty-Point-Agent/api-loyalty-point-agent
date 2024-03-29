package response

import (
	"api-loyalty-point-agent/businesses/stock_details"

	"api-loyalty-point-agent/businesses/transactions"

	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID             uint                 `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time            `json:"created_at"`
	UpdatedAt      time.Time            `json:"updated_at"`
	DeletedAt      gorm.DeletedAt       `json:"deleted_at"`
	Phone          string               `json:"phone"`
	StockID        uint                 `json:"stock_id"`
	StockDetailsID uint                 `json:"stock_details_id" validate:"required"`
	StockDetail    stock_details.Domain `json:"-"`
	Price          float64              `json:"price"`
	Product        string               `json:"product"`
	Payment_method string               `json:"payment_method"`
	Point          float64              `json:"point"`
	Status         string               `json:"status"`
	Description    string               `json:"description"`
	UserID         uint                 `json:"user_id"`
}

type TransactionInAdmin struct {
	ID             uint                 `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time            `json:"created_at"`
	UpdatedAt      time.Time            `json:"updated_at"`
	DeletedAt      gorm.DeletedAt       `json:"deleted_at"`
	Phone          string               `json:"phone"`
	StockID        uint                 `json:"stock_id"`
	StockDetailsID uint                 `json:"stock_details_id" validate:"required"`
	StockDetail    stock_details.Domain `json:"-"`
	Price          float64              `json:"price"`
	Product        string               `json:"product"`
	Payment_method string               `json:"payment_method"`
	Point          float64              `json:"point"`
	Status         string               `json:"status"`
	Description    string               `json:"description"`
	UserID         uint                 `json:"user_id"`
	UserName       string               `json:"user_name"`
	Member         string               `json:"member"`
	URL            string               `json:"url"`
}

func FromDomain(domain transactions.Domain) Transaction {
	return Transaction{
		ID:             domain.ID,
		Phone:          domain.Phone,
		Product:        domain.Product,
		StockDetailsID: domain.StockDetailsID,
		StockDetail:    domain.StockDetails,
		StockID:        domain.StockID,
		Price:          domain.Price,
		Payment_method: domain.Payment_method,
		Point:          domain.Point,
		Status:         domain.Status,
		Description:    domain.Description,
		UserID:         domain.UserID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
	}
}

func FromDomainInAdmin(domain transactions.Domain) TransactionInAdmin {
	return TransactionInAdmin{
		ID:             domain.ID,
		Phone:          domain.Phone,
		Product:        domain.Product,
		StockDetailsID: domain.StockDetailsID,
		StockDetail:    domain.StockDetails,
		StockID:        domain.StockID,
		Price:          domain.Price,
		Payment_method: domain.Payment_method,
		Point:          domain.Point,
		Status:         domain.Status,
		Description:    domain.Description,
		UserID:         domain.UserID,
		UserName:       domain.UserName,
		Member:         domain.Member,
		URL:            domain.URL,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
	}
}
