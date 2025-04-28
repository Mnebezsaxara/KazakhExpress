package usecase

import (
	"context"
	"errors"

	"auth-service/infrastructure/token"
	"auth-service/internal/domain"
	"auth-service/pkg/hash"
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
	_, err := uc.userRepo.GetUserByEmail(ctx, email)
	if err == nil {
		return errors.New("user already exists")
	}

	hashedPwd, err := hash.HashPassword(password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: hashedPwd,
		Role:     domain.RoleUser,
	}

	return uc.userRepo.CreateUser(ctx, user)
}

func (uc *AuthUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := uc.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !hash.CheckPasswordHash(user.Password, password) {
		return "", errors.New("invalid email or password")
	}

	return uc.jwtManager.Generate(user.Email, user.Role)
}

func (uc *AuthUsecase) GetProfile(ctx context.Context, email string) (*domain.User, error) {
	user, err := uc.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	// Never return the hashed password
	user.Password = ""
	return user, nil
}

func (uc *AuthUsecase) UpdateUser(ctx context.Context, email, newName, newPassword, newRole string) error {
	user, err := uc.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	if newPassword != "" {
		hashedPwd, err := hash.HashPassword(newPassword)
		if err != nil {
			return err
		}
		user.Password = hashedPwd
	}

	if newName != "" {
		user.Name = newName
	}

	if newRole != "" {
		user.Role = newRole
	}

	return uc.userRepo.UpdateUser(ctx, email, user)
}

func (uc *AuthUsecase) DeleteUser(ctx context.Context, email string) error {
	return uc.userRepo.DeleteUser(ctx, email)
}
