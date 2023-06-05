package user_details

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID               uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	TPoint           int64
	Member           string
	Gender           string
	Age              string
	Address          string
	Email            string
	PhoneNumber      string
	URL              string
	CountTransaction int64
	CountRedeem      int64
}
type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	Update(ctx context.Context, user_detailDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
}
type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	Update(ctx context.Context, user_detailDomain *Domain, id string) (Domain, error)
	Delete(ctx context.Context, id string) error
}
