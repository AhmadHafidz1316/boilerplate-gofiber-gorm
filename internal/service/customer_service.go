package service

import (
	"context"

	"github.com/go-playground/validator/v10"

	"github.com/AhmadHafidz1316/go-fiber-gorm/internal/domain"
)

type CustomerService interface {
	List(ctx context.Context) ([]domain.Customer, error)
	Get(ctx context.Context, id uint) (*domain.Customer, error)
	Create(ctx context.Context, in CreateCustomerInput) (*domain.Customer, error)
	Update(ctx context.Context, id uint, in UpdateCustomerInput) (*domain.Customer, error)
	Delete(ctx context.Context, id uint) error
}

type customerService struct {
	repo domain.CustomerRepository
	v    *validator.Validate
}

func NewCustomerService(repo domain.CustomerRepository) CustomerService {
	return &customerService{repo: repo, v: validator.New()}
}

type CreateCustomerInput struct {
	Name  string `json:"name" validate:"required,min=2"`
	Email string `json:"email" validate:"omitempty,email"`
}

type UpdateCustomerInput struct {
	Name  *string `json:"name" validate:"omitempty,min=2"`
	Email *string `json:"email" validate:"omitempty,email"`
}

func (s *customerService) List(ctx context.Context) ([]domain.Customer, error) {
	return s.repo.FindAll(ctx)
}

func (s *customerService) Get(ctx context.Context, id uint) (*domain.Customer, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *customerService) Create(ctx context.Context, in CreateCustomerInput) (*domain.Customer, error) {
	if err := s.v.Struct(in); err != nil {
		return nil, err
	}
	c := &domain.Customer{Name: in.Name, Email: in.Email}
	if err := s.repo.Create(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *customerService) Update(ctx context.Context, id uint, in UpdateCustomerInput) (*domain.Customer, error) {
	if err := s.v.Struct(in); err != nil {
		return nil, err
	}
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, nil
	}
	if in.Name != nil {
		existing.Name = *in.Name
	}
	if in.Email != nil {
		existing.Email = *in.Email
	}
	if err := s.repo.Update(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *customerService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
