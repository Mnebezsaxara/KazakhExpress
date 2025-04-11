package server

import (
	"fmt"
	"net"

	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/domain"
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
	grpcServer := grpc.NewServer()

	productServer := NewProductServer(productUsecase)
	categoryServer := NewCategoryServer(categoryUsecase, productUsecase)

	pb.RegisterProductServiceServer(grpcServer, productServer)
	pb.RegisterCategoryServiceServer(grpcServer, categoryServer)

	// Enable reflection for debugging
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

	fmt.Printf("gRPC server listening on port %d\n", s.port)
	if err := s.grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
} 