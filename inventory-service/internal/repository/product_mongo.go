package repository

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productMongo struct {
	collection *mongo.Collection
}

func NewProductMongo(db *mongo.Collection) domain.ProductRepository {
	return &productMongo{
		collection: db,
	}
}

func (r *productMongo) Create(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now

	oid := primitive.NewObjectID()
	product.ID = oid.Hex()

	_, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productMongo) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	var product domain.Product

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ErrProductNotFound
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrProductNotFound
		}
		return nil, err
	}

	return &product, nil
}

func (r *productMongo) List(ctx context.Context, filter domain.ProductFilter) ([]*domain.Product, int, error) {
	findOptions := options.Find()
	findOptions.SetSkip(int64((filter.Page - 1) * filter.Limit))
	findOptions.SetLimit(int64(filter.Limit))

	query := bson.M{}
	if filter.Category != "" {
		query["category"] = filter.Category
	}
	if filter.MinPrice > 0 {
		query["price"] = bson.M{"$gte": filter.MinPrice}
	}
	if filter.MaxPrice > 0 {
		if _, ok := query["price"]; ok {
			query["price"].(bson.M)["$lte"] = filter.MaxPrice
		} else {
			query["price"] = bson.M{"$lte": filter.MaxPrice}
		}
	}

	cursor, err := r.collection.Find(ctx, query, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var products []*domain.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, query)
	if err != nil {
		return nil, 0, err
	}

	return products, int(total), nil
}

func (r *productMongo) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	product.UpdatedAt = time.Now()

	oid, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		return nil, domain.ErrProductNotFound
	}

	result, err := r.collection.ReplaceOne(ctx, bson.M{"_id": oid}, product)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, domain.ErrProductNotFound
	}

	return product, nil
}

func (r *productMongo) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrProductNotFound
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return domain.ErrProductNotFound
	}

	return nil
}

func (r *productMongo) FindAll(filters map[string]interface{}, limit, offset int) ([]domain.Product, error) {
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
