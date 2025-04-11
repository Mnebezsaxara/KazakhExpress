package usecase

import "github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/domain"

type CategoryRepository interface {
    Insert(category *domain.Category) error
    FindAll() ([]domain.Category, error)
    FindByID(id string) (domain.Category, error)
    Delete(id string) error
}

type CategoryUsecase struct {
    repo CategoryRepository
}

func NewCategoryUsecase(r CategoryRepository) *CategoryUsecase {
    return &CategoryUsecase{repo: r}
}

func (uc *CategoryUsecase) CreateCategory(c *domain.Category) error {
    return uc.repo.Insert(c)
}

func (uc *CategoryUsecase) GetCategories() ([]domain.Category, error) {
    return uc.repo.FindAll()
}

func (uc *CategoryUsecase) GetCategoryByID(id string) (domain.Category, error) {
    return uc.repo.FindByID(id)
}

func (uc *CategoryUsecase) RemoveCategory(id string) error {
    return uc.repo.Delete(id)
}
