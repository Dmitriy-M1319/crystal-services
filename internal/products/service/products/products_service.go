package products

import (
	"context"

	"github.com/pkg/errors"
)

type ProductRepository interface {
	GetProducts(ctx context.Context) ([]Product, error)
	GetProductByID(ctx context.Context, id uint64) (Product, error)
	InsertProduct(ctx context.Context, product *Product) error
	UpdateProduct(ctx context.Context, id uint64, product *Product) error
	DeleteProduct(ctx context.Context, id uint64) error
}

type Service struct {
	repository ProductRepository
}

func NewService(r ProductRepository) *Service {
	return &Service{repository: r}
}

func (s Service) GetAllProducts(ctx context.Context, nameFilter *CompanyNameFilter, priceFilter *ClientPriceFilter) ([]Product, error) {
	ps, err := s.repository.GetProducts(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "repository.GetProducts")
	}

	if nameFilter != nil {
		ps = FilterByCompany(ps, nameFilter)
	}
	if priceFilter != nil {
		ps = FilterByPrice(ps, priceFilter)
	}

	return ps, nil
}

func (s Service) GetProduct(ctx context.Context, id uint64) (Product, error) {
	p, err := s.repository.GetProductByID(ctx, id)
	if err != nil {
		return Product{}, errors.Wrap(err, "repository.GetProductByID")
	}

	return p, nil
}

func (s Service) AddProduct(ctx context.Context, newProd *Product) error {
	err := s.repository.InsertProduct(ctx, newProd)
	if err != nil {
		return errors.Wrap(err, "repository.InsertProduct")
	}

	return nil
}

func (s Service) UpdateProduct(ctx context.Context, id uint64, prod *Product) error {
	err := s.repository.UpdateProduct(ctx, id, prod)
	if err != nil {
		return errors.Wrap(err, "repository.UpdateProduct")
	}

	return nil
}

func (s Service) DeleteProduct(ctx context.Context, id uint64) error {
	err := s.repository.DeleteProduct(ctx, id)
	if err != nil {
		return errors.Wrap(err, "repository.DeleteProduct")
	}

	return nil
}
