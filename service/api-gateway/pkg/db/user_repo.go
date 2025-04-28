package db

import (
	"context"

	"api-gateway/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateUser(ctx context.Context, email string, updated *domain.User) error
	DeleteUser(ctx context.Context, email string) error
}

type userRepo struct {
	collection *mongo.Collection
}

func NewUserRepository(col *mongo.Collection) UserRepository {
	return &userRepo{
		collection: col,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *domain.User) error {
	dbUser := &User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}

	_, err := r.collection.InsertOne(ctx, dbUser)
	return err
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var dbUser User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&dbUser)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       dbUser.ID.Hex(),
		Name:     dbUser.Name,
		Email:    dbUser.Email,
		Password: dbUser.Password,
		Role:     dbUser.Role,
	}, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, email string, updated *domain.User) error {
	dbUser := &User{
		Name:     updated.Name,
		Email:    updated.Email,
		Password: updated.Password,
		Role:     updated.Role,
	}

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"email": email},
		bson.M{"$set": dbUser},
	)
	return err
}

func (r *userRepo) DeleteUser(ctx context.Context, email string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"email": email})
	return err
} 