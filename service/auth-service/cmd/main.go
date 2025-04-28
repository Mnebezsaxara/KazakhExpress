package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"auth-service/config"
	"auth-service/infrastructure/db"
	"auth-service/infrastructure/token"
	handlerGrpc "auth-service/internal/grpc"    // gRPC handler
	handlerRest "auth-service/internal/handler" // REST handler
	"auth-service/internal/usecase"
	"auth-service/proto"
)

func main() {
	_ = godotenv.Load() // Load .env

	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	cfg := config.LoadConfig()

	// Connect to MongoDB
	client, database, err := db.ConnectMongoDB(cfg.MongoURI, cfg.DatabaseName)
	if err != nil {
		log.Fatal("Mongo connection failed:", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := client.Disconnect(ctx); err != nil {
			log.Println("Error disconnecting Mongo:", err)
		}
	}()

	// Set up dependencies
	userCollection := database.Collection("users")
	userRepo := db.NewUserRepository(userCollection)
	jwtManager := token.NewJWTManager(cfg.JWTSecret, cfg.JWTDuration)
	authUC := usecase.NewAuthUsecase(userRepo, jwtManager)

	// REST setup
	router := gin.New()
	router.Use(gin.Recovery())
	
	// CORS config
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(corsConfig))

	// Set trusted proxies
	router.SetTrustedProxies([]string{"127.0.0.1"})

	authHandler := handlerRest.NewAuthHandler(authUC)
	authHandler.RegisterRoutes(router)

	// ✅ Start REST in a goroutine
	go func() {
		log.Printf("🔵 REST API running on port %s...\n", cfg.HTTPPort)
		if err := router.Run(":" + cfg.HTTPPort); err != nil {
			log.Fatal("Failed to start REST server:", err)
		}
	}()

	// ✅ Start gRPC server
	grpcAddress := fmt.Sprintf(":%s", cfg.Port)
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("Failed to listen on gRPC port %s: %v", grpcAddress, err)
	}

	grpcServer := grpc.NewServer()
	grpcHandler := handlerGrpc.NewGrpcHandler(authUC)
	proto.RegisterUserServiceServer(grpcServer, grpcHandler)

	log.Printf("🟢 gRPC server running on port %s...\n", cfg.Port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
