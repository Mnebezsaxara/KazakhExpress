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

type ProductServer struct {
	pb.UnimplementedProductServiceServer
	productUsecase domain.ProductUsecase
}

func NewProductServer(productUsecase domain.ProductUsecase) *ProductServer {
	return &ProductServer{
		productUsecase: productUsecase,
	}
}

func (s *ProductServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
	product := &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		ImageURL:    req.ImageUrl,
		Category:    req.Category,
		Stock:       int(req.Stock),
	}

	result, err := s.productUsecase.Create(ctx, product)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create product: %v", err)
	}

	return convertDomainToProtoProduct(result), nil
}

func (s *ProductServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	product, err := s.productUsecase.GetByID(ctx, req.Id)
	if err != nil {
		if err == domain.ErrProductNotFound {
			return nil, status.Errorf(codes.NotFound, "product not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get product: %v", err)
	}

	return convertDomainToProtoProduct(product), nil
}

func (s *ProductServer) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	filter := domain.ProductFilter{
		Category: req.Category,
		MinPrice: req.MinPrice,
		MaxPrice: req.MaxPrice,
		Page:     int(req.Page),
		Limit:    int(req.Limit),
	}

	products, total, err := s.productUsecase.List(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list products: %v", err)
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

func (s *ProductServer) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error) {
	product := &domain.Product{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		ImageURL:    req.ImageUrl,
		Category:    req.Category,
		Stock:       int(req.Stock),
	}

	result, err := s.productUsecase.Update(ctx, product)
	if err != nil {
		if err == domain.ErrProductNotFound {
			return nil, status.Errorf(codes.NotFound, "product not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update product: %v", err)
	}

	return convertDomainToProtoProduct(result), nil
}

func (s *ProductServer) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*emptypb.Empty, error) {
	err := s.productUsecase.Delete(ctx, req.Id)
	if err != nil {
		if err == domain.ErrProductNotFound {
			return nil, status.Errorf(codes.NotFound, "product not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete product: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func convertDomainToProtoProduct(product *domain.Product) *pb.Product {
	return &pb.Product{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		ImageUrl:    product.ImageURL,
		Category:    product.Category,
		Stock:       int32(product.Stock),
		CreatedAt:   timestamppb.New(product.CreatedAt),
		UpdatedAt:   timestamppb.New(product.UpdatedAt),
	}
} 