package usecase

import (
	"context"
	"errors"

	"api-gateway/internal/domain"
	"api-gateway/pkg/hash"
	"api-gateway/pkg/token"
)

type AuthUsecase struct {
	userRepo   domain.UserRepository
	jwtManager *token.JWTManager
}

func NewAuthUsecase(repo domain.UserRepository, jwt *token.JWTManager) *AuthUsecase {
	return &AuthUsecase{
		userRepo:   repo,
		jwtManager: jwt,
	}
}

func (uc *AuthUsecase) SignUp(ctx context.Context, name, email, password string) error {
	// Check if user already exists
	existing, err := uc.userRepo.GetUserByEmail(ctx, email)
	if err == nil && existing != nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		return err
	}

	// Create new user
	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Role:     "user", // default role
	}

	return uc.userRepo.CreateUser(ctx, user)
}

func (uc *AuthUsecase) Login(ctx context.Context, email, password string) (string, error) {
	// Get user
	user, err := uc.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Check password
	if !hash.CheckPasswordHash(user.Password, password) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := uc.jwtManager.Generate(user.Email, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *AuthUsecase) GetProfile(ctx context.Context, email string) (*domain.User, error) {
	return uc.userRepo.GetUserByEmail(ctx, email)
}

func (uc *AuthUsecase) UpdateUser(ctx context.Context, email, newName, newPassword, newRole string) error {
	// Get existing user
	user, err := uc.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	// Update fields if provided
	if newName != "" {
		user.Name = newName
	}
	if newPassword != "" {
		hashedPassword, err := hash.HashPassword(newPassword)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}
	if newRole != "" {
		user.Role = newRole
	}

	return uc.userRepo.UpdateUser(ctx, email, user)
}

func (uc *AuthUsecase) DeleteUser(ctx context.Context, email string) error {
	return uc.userRepo.DeleteUser(ctx, email)
}
