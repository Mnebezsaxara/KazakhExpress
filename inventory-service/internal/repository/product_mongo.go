package repository

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/Mnebezsaxara/inventory-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type ProductMongo struct {
    collection *mongo.Collection
}

func NewProductMongo(col *mongo.Collection) *ProductMongo {
    return &ProductMongo{collection: col}
}

func (r *ProductMongo) Insert(product *domain.Product) error {
    // Генерируем новый ObjectID для продукта
    product.ID = primitive.NewObjectID()
    
    // Вставляем продукт в базу данных
    result, err := r.collection.InsertOne(context.TODO(), product)
    if err != nil {
        return err
    }
    
    // Проверяем, что ID был успешно установлен
    if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
        product.ID = oid
    }
    
    return nil
}

func (r *ProductMongo) FindAll(filters map[string]interface{}, limit, offset int) ([]domain.Product, error) {
    var products []domain.Product
    options := options.Find()

    // Создаем фильтр
    filter := bson.M{}

    // Логируем входящие параметры
    log.Printf("FindAll called with filters: %+v, limit: %d, offset: %d", filters, limit, offset)

    // Применяем фильтры
    if category, ok := filters["category"]; ok && category != "" {
        filter["category"] = category
        log.Printf("Applied category filter: %s", category)
    }

    // Обработка minPrice и maxPrice
    if minPrice, ok := filters["minPrice"]; ok && minPrice != "" {
        minPriceFloat, err := strconv.ParseFloat(minPrice.(string), 64)
        if err == nil {
            if filter["price"] == nil {
                filter["price"] = bson.M{}
            }
            filter["price"].(bson.M)["$gte"] = minPriceFloat
            log.Printf("Applied minPrice filter: %f", minPriceFloat)
        }
    }

    if maxPrice, ok := filters["maxPrice"]; ok && maxPrice != "" {
        maxPriceFloat, err := strconv.ParseFloat(maxPrice.(string), 64)
        if err == nil {
            if filter["price"] == nil {
                filter["price"] = bson.M{}
            }
            filter["price"].(bson.M)["$lte"] = maxPriceFloat
            log.Printf("Applied maxPrice filter: %f", maxPriceFloat)
        }
    }

    // Применяем пагинацию
    if limit > 0 {
        options.SetLimit(int64(limit))
    }
    if offset > 0 {
        options.SetSkip(int64(offset))
    }

    // Логируем финальный фильтр
    log.Printf("Final filter: %+v", filter)

    // Выполняем запрос к базе данных
    cursor, err := r.collection.Find(context.TODO(), filter, options)
    if err != nil {
        log.Printf("Error executing find query: %v", err)
        return nil, err
    }
    defer cursor.Close(context.TODO())

    // Декодируем результаты
    for cursor.Next(context.TODO()) {
        var product domain.Product
        if err := cursor.Decode(&product); err != nil {
            log.Printf("Error decoding product: %v", err)
            return nil, err
        }
        products = append(products, product)
    }

    if err := cursor.Err(); err != nil {
        log.Printf("Cursor error: %v", err)
        return nil, err
    }

    // Логируем количество найденных продуктов
    log.Printf("Found %d products", len(products))

    return products, nil
}


func (r *ProductMongo) FindByID(id string) (*domain.Product, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }
    var product domain.Product
    err = r.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&product)
    if err != nil {
        return nil, err
    }
    return &product, nil
}

func (r *ProductMongo) Update(id string, update bson.M) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    result, err := r.collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": update})
    if err != nil {
        return err
    }
    if result.MatchedCount == 0 {
        return errors.New("продукт не найден")
    }
    return nil
}

func (r *ProductMongo) Delete(id string) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    result, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
    if err != nil {
        return err
    }
    if result.DeletedCount == 0 {
        return errors.New("продукт не найден")
    }
    return nil
}
