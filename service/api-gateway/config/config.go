package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
    Port              string
    MongoURI          string
    DatabaseName      string
    JWTSecret         string
    JWTDuration       time.Duration
    HTTPPort          string
    AuthServiceURL    string
    InventoryServiceURL string
    OrderServiceURL    string
}

func LoadConfig() *Config {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI not set")
    }

    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "authdb"
    }

    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        log.Fatal("JWT_SECRET not set")
    }

    httpPort := os.Getenv("HTTP_PORT")
    if httpPort == "" {
        httpPort = "8082"
    }

    authURL := os.Getenv("AUTH_SERVICE_URL")
    if authURL == "" {
        authURL = "auth-service:50051"
    }

    inventoryURL := os.Getenv("INVENTORY_SERVICE_URL")
    if inventoryURL == "" {
        inventoryURL = "inventory-service:50052"
    }

    orderURL := os.Getenv("ORDER_SERVICE_URL")
    if orderURL == "" {
        orderURL = "order-service:50053"
    }

    return &Config{
        Port:               port,
        MongoURI:          mongoURI,
        DatabaseName:      dbName,
        JWTSecret:         jwtSecret,
        JWTDuration:       time.Hour * 24,
        HTTPPort:          httpPort,
        AuthServiceURL:    authURL,
        InventoryServiceURL: inventoryURL,
        OrderServiceURL:    orderURL,
    }
}