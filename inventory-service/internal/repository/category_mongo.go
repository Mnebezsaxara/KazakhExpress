package repository

import (
	"context"

	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryMongo struct {
	collection *mongo.Collection
}

func NewCategoryMongo(col *mongo.Collection) domain.CategoryRepository {
	return &CategoryMongo{collection: col}
}

func (r *CategoryMongo) Create(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	oid := primitive.NewObjectID()
	category.ID = oid.Hex()

	_, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryMongo) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	var category domain.Category
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ErrCategoryNotFound
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrCategoryNotFound
		}
		return nil, err
	}

	return &category, nil
}

func (r *CategoryMongo) List(ctx context.Context) ([]*domain.Category, int, error) {
	var categories []*domain.Category
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &categories); err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return categories, int(total), nil
}

func (r *CategoryMongo) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrCategoryNotFound
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return domain.ErrCategoryNotFound
	}

	return nil
}

func (r *CategoryMongo) UpdateProductCount(ctx context.Context, id string, delta int) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrCategoryNotFound
	}

	update := bson.M{
		"$inc": bson.M{"product_count": delta},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return domain.ErrCategoryNotFound
	}

	return nil
}
