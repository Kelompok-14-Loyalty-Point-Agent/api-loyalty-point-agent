package redeems

// penyebab Error 1452 (23000): Cannot add or update a child row: a foreign key constraint fails (`api_loyalty_point_agent_db`.`redeems`, CONSTRAINT `fk_redeems_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`))
import (
	"api-loyalty-point-agent/businesses/redeems"

	"time"
)

type Redeem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"create_at"`
	Phone     string    `json:"phone"`
	Cost      uint      `json:"cost"`
	//ini
	// UserID uint `json:"user_id" gorm:"foreignKey:UserID"`
}

func (rec *Redeem) ToDomain() redeems.Domain {
	return redeems.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		Phone:     rec.Phone,
		Cost:      rec.Cost,
		// UserID:    rec.UserID,
	}
}

func FromDomain(domain *redeems.Domain) *Redeem {
	return &Redeem{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		Phone:     domain.Phone,
		Cost:      domain.Cost,
		// UserID:    domain.UserID,
	}
}
