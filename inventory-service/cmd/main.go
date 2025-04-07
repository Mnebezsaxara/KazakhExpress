package main

import (
	"log"

	"github.com/Mnebezsaxara/inventory-service/config"
	"github.com/Mnebezsaxara/inventory-service/internal/handler"
	"github.com/Mnebezsaxara/inventory-service/internal/repository"
	"github.com/Mnebezsaxara/inventory-service/internal/usecase"
	"github.com/gin-contrib/cors" // Добавляем CORS middleware
	"github.com/gin-gonic/gin"
)

func main() {
    col := config.ConnectMongo()

    // Репозиторий и usecase для продуктов
    productRepo := repository.NewProductMongo(col)
    productUC := usecase.NewProductUsecase(productRepo)
    productHandler := handler.NewProductHandler(productUC)

    // Репозиторий и usecase для категорий
    categoryRepo := repository.NewCategoryMongo(col)
    categoryUC := usecase.NewCategoryUsecase(categoryRepo)
    categoryHandler := handler.NewCategoryHandler(categoryUC)

    r := gin.Default()

    // Добавляем CORS middleware
    r.Use(cors.Default())

    // Маршруты для продуктов
    r.POST("/products", productHandler.Create)
    r.GET("/products", productHandler.List)
    r.GET("/products/:id", productHandler.GetByID)
    r.PATCH("/products/:id", productHandler.Update)
    r.DELETE("/products/:id", productHandler.Delete)

    // Маршруты для категорий
    r.POST("/categories", categoryHandler.Create)
    r.GET("/categories", categoryHandler.List)
    r.GET("/categories/:id", categoryHandler.GetByID)
    r.DELETE("/categories/:id", categoryHandler.Delete)

    log.Println("Inventory Service запущен на http://localhost:8081")
    r.Run(":8081")
}
