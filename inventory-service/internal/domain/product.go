package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type Product struct {
	ID          string    `bson:"_id,omitempty"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	Price       float64   `bson:"price"`
	ImageURL    string    `bson:"image_url"`
	Category    string    `bson:"category"`
	Stock       int       `bson:"stock"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

type ProductFilter struct {
	Category string
	MinPrice float64
	MaxPrice float64
	Page     int
	Limit    int
}

type ProductRepository interface {
	Create(ctx context.Context, product *Product) (*Product, error)
	GetByID(ctx context.Context, id string) (*Product, error)
	List(ctx context.Context, filter ProductFilter) ([]*Product, int, error)
	Update(ctx context.Context, product *Product) (*Product, error)
	Delete(ctx context.Context, id string) error
}

type ProductUsecase interface {
	Create(ctx context.Context, product *Product) (*Product, error)
	GetByID(ctx context.Context, id string) (*Product, error)
	List(ctx context.Context, filter ProductFilter) ([]*Product, int, error)
	Update(ctx context.Context, product *Product) (*Product, error)
	Delete(ctx context.Context, id string) error
}