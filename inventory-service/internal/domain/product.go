package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name        string             `json:"name"`
    Description string             `json:"description"`
    Price       float64            `json:"price"`
    ImageURL    string             `json:"image_url"`
    Category    string             `json:"category"`
	Stock       int                `json:"stock"`  // Количество на складе
}