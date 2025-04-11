package server

import (
	"context"

	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/domain"
	pb "github.com/Mnebezsaxara/KazakhExpress/inventory-service/proto/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CategoryServer struct {
	pb.UnimplementedCategoryServiceServer
	categoryUsecase domain.CategoryUsecase
	productUsecase  domain.ProductUsecase
}

func NewCategoryServer(categoryUsecase domain.CategoryUsecase, productUsecase domain.ProductUsecase) *CategoryServer {
	return &CategoryServer{
		categoryUsecase: categoryUsecase,
		productUsecase:  productUsecase,
	}
}

func (s *CategoryServer) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.Category, error) {
	category := &domain.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	result, err := s.categoryUsecase.Create(ctx, category)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create category: %v", err)
	}

	return convertDomainToProtoCategory(result), nil
}

func (s *CategoryServer) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := s.categoryUsecase.GetByID(ctx, req.Id)
	if err != nil {
		if err == domain.ErrCategoryNotFound {
			return nil, status.Errorf(codes.NotFound, "category not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get category: %v", err)
	}

	return convertDomainToProtoCategory(category), nil
}

func (s *CategoryServer) ListCategories(ctx context.Context, _ *emptypb.Empty) (*pb.ListCategoriesResponse, error) {
	categories, total, err := s.categoryUsecase.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list categories: %v", err)
	}

	protoCategories := make([]*pb.Category, len(categories))
	for i, category := range categories {
		protoCategories[i] = convertDomainToProtoCategory(category)
	}

	return &pb.ListCategoriesResponse{
		Categories: protoCategories,
		Total:     int32(total),
	}, nil
}

func (s *CategoryServer) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*emptypb.Empty, error) {
	err := s.categoryUsecase.Delete(ctx, req.Id)
	if err != nil {
		if err == domain.ErrCategoryNotFound {
			return nil, status.Errorf(codes.NotFound, "category not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete category: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *CategoryServer) GetProductsByCategory(ctx context.Context, req *pb.GetProductsByCategoryRequest) (*pb.ListProductsResponse, error) {
	filter := domain.ProductFilter{
		Category: req.CategoryId,
		Page:     int(req.Page),
		Limit:    int(req.Limit),
	}

	products, total, err := s.productUsecase.List(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get products by category: %v", err)
	}

	protoProducts := make([]*pb.Product, len(products))
	for i, product := range products {
		protoProducts[i] = convertDomainToProtoProduct(product)
	}

	return &pb.ListProductsResponse{
		Products: protoProducts,
		Total:    int32(total),
	}, nil
}

func convertDomainToProtoCategory(category *domain.Category) *pb.Category {
	return &pb.Category{
		Id:           category.ID,
		Name:         category.Name,
		Description:  category.Description,
		ProductCount: int32(category.ProductCount),
		CreatedAt:    timestamppb.New(category.CreatedAt),
		UpdatedAt:    timestamppb.New(category.UpdatedAt),
	}
} 