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

func NewCategoryMongo(col *mongo.Collection) *CategoryMongo {
    return &CategoryMongo{collection: col}
}

// InsertCategory добавляет новую категорию
func (r *CategoryMongo) Insert(category *domain.Category) error {
    _, err := r.collection.InsertOne(context.TODO(), category)
    return err
}

// FindAllCategories возвращает все категории
func (r *CategoryMongo) FindAll() ([]domain.Category, error) {
    var categories []domain.Category
    cursor, err := r.collection.Find(context.TODO(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())
    for cursor.Next(context.TODO()) {
        var category domain.Category
        if err := cursor.Decode(&category); err != nil {
            return nil, err
        }
        categories = append(categories, category)
    }
    return categories, nil
}

// FindCategoryByID ищет категорию по ID
func (r *CategoryMongo) FindByID(id string) (domain.Category, error) {
    var category domain.Category
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return category, err
    }
    err = r.collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&category)
    return category, err
}

// DeleteCategory удаляет категорию по ID
func (r *CategoryMongo) Delete(id string) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
    return err
}
