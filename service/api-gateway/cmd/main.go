package main

import (
	"context"
	"fmt"
	"log"

	"api-gateway/config"
	"api-gateway/internal/client"
	"api-gateway/internal/handler"
	restHandler "api-gateway/internal/handler"
	"api-gateway/internal/usecase"
	"api-gateway/pkg/db"
	"api-gateway/pkg/token"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to MongoDB
	mongoClient, mongoDB, err := db.ConnectMongoDB(cfg.MongoURI, cfg.DatabaseName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoClient.Disconnect(context.Background())

	// Initialize JWT manager
	jwtManager := token.NewJWTManager(cfg.JWTSecret, cfg.JWTDuration)

	// Initialize repositories
	userCollection := mongoDB.Collection("users")
	userRepo := db.NewUserRepository(userCollection)

	// Initialize usecases
	authUC := usecase.NewAuthUsecase(userRepo, jwtManager)

	// Initialize gRPC clients
	invConn, err := grpc.Dial(cfg.InventoryServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to inventory service: %v", err)
	}
	defer invConn.Close()

	inventoryClient := client.NewInventoryClient(invConn)
	defer inventoryClient.Close()

	// Initialize REST handlers
	router := gin.New()
	router.Use(gin.Recovery())

	// Настройка CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://127.0.0.1:5501"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))

	// Set trusted proxies
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Initialize handlers
	authHandler := restHandler.NewAuthHandler(authUC)
	authHandler.RegisterRoutes(router)

	// Initialize inventory handler
	inventoryHandler := handler.NewHandler(inventoryClient)
	router.GET("/categories", inventoryHandler.HandleCategories)
	router.GET("/products", inventoryHandler.HandleProducts)
	router.POST("/orders", inventoryHandler.HandleOrders)
	router.GET("/orders", inventoryHandler.HandleGetOrders)
	router.DELETE("/orders/:id", inventoryHandler.HandleCancelOrder)

	// Serve static files (HTML, CSS, JS)
	// Serve static files
	router.Static("/static", "./public")

	// Serve index.html by default
	router.GET("/", func(c *gin.Context) {
		c.File("./public/html/index.html")
	})

	// Start HTTP server
	httpAddress := fmt.Sprintf(":%s", cfg.HTTPPort)
	log.Printf("Starting HTTP server on %s", httpAddress)
	if err := router.Run(httpAddress); err != nil {
		log.Fatalf("Failed to run HTTP server: %v", err)
	}
}
