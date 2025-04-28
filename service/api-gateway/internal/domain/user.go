package domain

import "context"

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

type User struct {
	ID       string `bson:"_id,omitempty"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"` // hashed password
	Role     string `bson:"role"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, email string, updated *User) error
	DeleteUser(ctx context.Context, email string) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}
