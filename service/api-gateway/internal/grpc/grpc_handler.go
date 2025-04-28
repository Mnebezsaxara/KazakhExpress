package handler

import (
	"context"

	"api-gateway/internal/usecase"
	"api-gateway/proto"
)

type GrpcHandler struct {
	proto.UnimplementedUserServiceServer
	AuthUC *usecase.AuthUsecase
}

func NewGrpcHandler(authUC *usecase.AuthUsecase) *GrpcHandler {
	return &GrpcHandler{
		AuthUC: authUC,
	}
}

func (h *GrpcHandler) RegisterUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	err := h.AuthUC.SignUp(ctx, req.Name, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &proto.UserResponse{Message: "User registered successfully"}, nil
}

func (h *GrpcHandler) AuthenticateUser(ctx context.Context, req *proto.AuthRequest) (*proto.AuthResponse, error) {
	token, err := h.AuthUC.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &proto.AuthResponse{Token: token}, nil
}

func (h *GrpcHandler) GetUserProfile(ctx context.Context, req *proto.UserID) (*proto.UserProfile, error) {
	user, err := h.AuthUC.GetProfile(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	return &proto.UserProfile{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (h *GrpcHandler) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UserResponse, error) {
	err := h.AuthUC.UpdateUser(ctx, req.Email, req.Name, req.Password, req.Role)
	if err != nil {
		return nil, err
	}
	return &proto.UserResponse{Message: "User updated successfully"}, nil
}

func (h *GrpcHandler) DeleteUser(ctx context.Context, req *proto.UserID) (*proto.DeleteResponse, error) {
	err := h.AuthUC.DeleteUser(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	return &proto.DeleteResponse{Message: "User deleted successfully"}, nil
}
