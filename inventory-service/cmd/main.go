package main

import (
	"log"

	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/config"
	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/grpc/server"
	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/handler"
	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/repository"
	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	col := config.ConnectMongo()

	// Repositories
	productRepo := repository.NewProductMongo(col)
	categoryRepo := repository.NewCategoryMongo(col)

	// Use cases
	productUC := usecase.NewProductUsecase(productRepo)
	categoryUC := usecase.NewCategoryUsecase(categoryRepo)

	// HTTP handlers
	productHandler := handler.NewProductHandler(productUC)
	categoryHandler := handler.NewCategoryHandler(categoryUC)

	// gRPC server
	grpcServer := server.NewServer(productUC, categoryUC, 50051)
	go func() {
		if err := grpcServer.Start(); err != nil {
			log.Fatalf("failed to start gRPC server: %v", err)
		}
	}()

	// HTTP server
	r := gin.Default()
	r.Use(cors.Default())

	// Product routes
	r.POST("/products", productHandler.Create)
	r.GET("/products", productHandler.List)
	r.GET("/products/:id", productHandler.GetByID)
	r.PATCH("/products/:id", productHandler.Update)
	r.DELETE("/products/:id", productHandler.Delete)

	// Category routes
	r.POST("/categories", categoryHandler.Create)
	r.GET("/categories", categoryHandler.List)
	r.GET("/categories/:id", categoryHandler.GetByID)
	r.DELETE("/categories/:id", categoryHandler.Delete)

	log.Println("Inventory Service запущен на http://localhost:8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
