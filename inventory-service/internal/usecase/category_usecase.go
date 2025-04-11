package usecase

import (
	"context"

	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/domain"
)

type CategoryUsecase struct {
	repo domain.CategoryRepository
}

func NewCategoryUsecase(repo domain.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{
		repo: repo,
	}
}

func (u *CategoryUsecase) Create(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	return u.repo.Create(ctx, category)
}

func (u *CategoryUsecase) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *CategoryUsecase) List(ctx context.Context) ([]*domain.Category, int, error) {
	return u.repo.List(ctx)
}

func (u *CategoryUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
