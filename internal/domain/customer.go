package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"size:100;not null"`
	Email     string         `json:"email" gorm:"size:120;uniqueIndex"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]Customer, error)
	FindByID(ctx context.Context, id uint) (*Customer, error)
	Create(ctx context.Context, c *Customer) error
	Update(ctx context.Context, c *Customer) error
	Delete(ctx context.Context, id uint) error
}
