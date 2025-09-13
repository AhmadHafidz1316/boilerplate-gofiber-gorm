package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/AhmadHafidz1316/go-fiber-gorm/internal/domain"
)

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) domain.CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) FindAll(ctx context.Context) ([]domain.Customer, error) {
	var items []domain.Customer
	err := r.db.WithContext(ctx).Find(&items).Error
	return items, err
}

func (r *customerRepository) FindByID(ctx context.Context, id uint) (*domain.Customer, error) {
	var item domain.Customer
	err := r.db.WithContext(ctx).First(&item, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &item, err
}

func (r *customerRepository) Create(ctx context.Context, c *domain.Customer) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *customerRepository) Update(ctx context.Context, c *domain.Customer) error {
	return r.db.WithContext(ctx).Save(c).Error
}

func (r *customerRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Customer{}, id).Error
}
