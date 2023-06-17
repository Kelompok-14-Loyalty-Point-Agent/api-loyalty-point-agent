package stocks

import (
	"api-loyalty-point-agent/businesses/stocks"
	"api-loyalty-point-agent/drivers/mysql/providers"
	"api-loyalty-point-agent/drivers/mysql/stock_details"
	"api-loyalty-point-agent/drivers/mysql/stock_transactions"

	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID                uint                                  `json:"id" gorm:"primaryKey"`
	CreatedAt         time.Time                             `json:"created_at"`
	UpdatedAt         time.Time                             `json:"updated_at"`
	DeletedAt         gorm.DeletedAt                        `json:"deleted_at" gorm:"index"`
	Type              string                                `json:"type" gorm:"type:enum('credit', 'data')"`
	TotalStock        float64                               `json:"total_stock"`
	LastTopUp         time.Time                             `json:"last_top_up" gorm:"type:datetime;default:null"`
	ProviderID        uint                                  `json:"provider_id"`
	Provider          providers.Provider                    `json:"-" gorm:"foreignKey:ProviderID"`
	StockDetails      []stock_details.StockDetail           `json:"-" gorm:"foreignKey:StockID"`
	StockTransactions []stock_transactions.StockTransaction `json:"-" gorm:"foreignKey:StockID"`
}

func (rec *Stock) ToDomain() stocks.Domain {
	return stocks.Domain{
		ID:         rec.ID,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
		DeletedAt:  rec.DeletedAt,
		Type:       rec.Type,
		TotalStock: rec.TotalStock,
		LastTopUp:  rec.LastTopUp,
		ProviderID: rec.ProviderID,
		Provider:   rec.Provider.ToDomain(),
	}
}

func FromDomain(domain *stocks.Domain) *Stock {
	return &Stock{
		ID:         domain.ID,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:  domain.DeletedAt,
		Type:       domain.Type,
		TotalStock: domain.TotalStock,
		LastTopUp:  domain.LastTopUp,
		Provider:   *providers.FromDomain(&domain.Provider),
		ProviderID: domain.ProviderID,
	}
}
