package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

type Category struct {
	ID           string    `bson:"_id,omitempty"`
	Name         string    `bson:"name"`
	Description  string    `bson:"description"`
	ProductCount int       `bson:"product_count"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

type CategoryRepository interface {
	Create(ctx context.Context, category *Category) (*Category, error)
	GetByID(ctx context.Context, id string) (*Category, error)
	List(ctx context.Context) ([]*Category, int, error)
	Delete(ctx context.Context, id string) error
	UpdateProductCount(ctx context.Context, id string, delta int) error
}

type CategoryUsecase interface {
	Create(ctx context.Context, category *Category) (*Category, error)
	GetByID(ctx context.Context, id string) (*Category, error)
	List(ctx context.Context) ([]*Category, int, error)
	Delete(ctx context.Context, id string) error
}
