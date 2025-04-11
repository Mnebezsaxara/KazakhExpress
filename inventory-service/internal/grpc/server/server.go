package server

import (
	"fmt"
	"net"

	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/domain"
	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/logger"
	pb "github.com/Mnebezsaxara/KazakhExpress/inventory-service/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
	port       int
}

func NewServer(
	productUsecase domain.ProductUsecase,
	categoryUsecase domain.CategoryUsecase,
	port int,
) *Server {
	// Создаем gRPC сервер с логированием
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(logger.UnaryServerInterceptor()),
	)

	productServer := NewProductServer(productUsecase)
	categoryServer := NewCategoryServer(categoryUsecase, productUsecase)

	pb.RegisterProductServiceServer(grpcServer, productServer)
	pb.RegisterCategoryServiceServer(grpcServer, categoryServer)

	// Включаем reflection для отладки
	reflection.Register(grpcServer)

	return &Server{
		grpcServer: grpcServer,
		port:       port,
	}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	logger.LogInfo("gRPC server listening on port %d", s.port)
	if err := s.grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func (s *Server) Stop() {
	logger.LogInfo("Stopping gRPC server...")
	s.grpcServer.GracefulStop()
} 