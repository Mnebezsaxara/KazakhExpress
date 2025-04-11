package usecase

import (
	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type ProductRepository interface {
    Insert(product *domain.Product) error
    FindAll(filters map[string]interface{}, limit, offset int) ([]domain.Product, error)
    FindByID(id string) (*domain.Product, error)
    Update(id string, update bson.M) error
    Delete(id string) error
}

type ProductUsecase struct {
    repo ProductRepository
}

func NewProductUsecase(r ProductRepository) *ProductUsecase {
    return &ProductUsecase{repo: r}
}

func (uc *ProductUsecase) CreateProduct(p *domain.Product) error {
    return uc.repo.Insert(p)
}

func (uc *ProductUsecase) ListProducts() ([]domain.Product, error) {
    filters := make(map[string]interface{})
    return uc.repo.FindAll(filters, 10, 0) // Пример с пагинацией
}

func (uc *ProductUsecase) GetProductByID(id string) (*domain.Product, error) {
    return uc.repo.FindByID(id)
}

func (uc *ProductUsecase) UpdateProduct(id string, update bson.M) error {
    return uc.repo.Update(id, update)
}

func (uc *ProductUsecase) DeleteProduct(id string) error {
    return uc.repo.Delete(id)
}

func (uc *ProductUsecase) GetProducts(filters map[string]interface{}, limit, offset int) ([]domain.Product, error) {
    return uc.repo.FindAll(filters, limit, offset)
}
