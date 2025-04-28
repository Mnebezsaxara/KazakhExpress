package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
    Port         string // gRPC port
    HTTPPort     string // REST port
    MongoURI     string
    DatabaseName string
    JWTSecret    string
    JWTDuration  time.Duration
}

func LoadConfig() *Config {
    port := os.Getenv("PORT")
    if port == "" {
        port = "50051"
    }

    httpPort := os.Getenv("HTTP_PORT")
    if httpPort == "" {
        httpPort = "8081" // Different port for REST
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

    return &Config{
        Port:         port,
        HTTPPort:     httpPort,
        MongoURI:     mongoURI,
        DatabaseName: dbName,
        JWTSecret:    jwtSecret,
        JWTDuration:  time.Hour * 24,
    }
}