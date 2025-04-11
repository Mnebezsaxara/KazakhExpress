package usecase

import (
	"context"

	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/domain"
)

type ProductUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(repo domain.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		repo: repo,
	}
}

func (u *ProductUsecase) Create(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	return u.repo.Create(ctx, product)
}

func (u *ProductUsecase) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *ProductUsecase) List(ctx context.Context, filter domain.ProductFilter) ([]*domain.Product, int, error) {
	return u.repo.List(ctx, filter)
}

func (u *ProductUsecase) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	return u.repo.Update(ctx, product)
}

func (u *ProductUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
