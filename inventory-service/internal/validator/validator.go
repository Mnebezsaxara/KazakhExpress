package validator

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidateProduct проверяет корректность данных продукта
func ValidateProduct(name, description string, price float64, stock int) error {
	if name == "" {
		return status.Error(codes.InvalidArgument, "name is required")
	}
	if len(name) > 100 {
		return status.Error(codes.InvalidArgument, "name is too long (max 100 characters)")
	}
	if len(description) > 1000 {
		return status.Error(codes.InvalidArgument, "description is too long (max 1000 characters)")
	}
	if price < 0 {
		return status.Error(codes.InvalidArgument, "price cannot be negative")
	}
	if stock < 0 {
		return status.Error(codes.InvalidArgument, "stock cannot be negative")
	}
	return nil
}

// ValidateCategory проверяет корректность данных категории
func ValidateCategory(name, description string) error {
	if name == "" {
		return status.Error(codes.InvalidArgument, "name is required")
	}
	if len(name) > 50 {
		return status.Error(codes.InvalidArgument, "name is too long (max 50 characters)")
	}
	if len(description) > 500 {
		return status.Error(codes.InvalidArgument, "description is too long (max 500 characters)")
	}
	return nil
}

// ValidateID проверяет корректность ID
func ValidateID(id string) error {
	if id == "" {
		return status.Error(codes.InvalidArgument, "id is required")
	}
	if len(id) != 24 {
		return status.Error(codes.InvalidArgument, "invalid id format")
	}
	return nil
}

// ValidatePagination проверяет корректность параметров пагинации
func ValidatePagination(page, limit int32) error {
	if page < 1 {
		return status.Error(codes.InvalidArgument, "page must be greater than 0")
	}
	if limit < 1 {
		return status.Error(codes.InvalidArgument, "limit must be greater than 0")
	}
	if limit > 100 {
		return status.Error(codes.InvalidArgument, "limit cannot exceed 100")
	}
	return nil
} 