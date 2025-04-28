package config

import (
    "log"
    "os"
)

type Config struct {
    MongoURI           string
    DatabaseName       string
    GRPCPort          string
    InventoryServiceURL string
}

func LoadConfig() *Config {
    grpcPort := os.Getenv("GRPC_PORT")
    if grpcPort == "" {
        grpcPort = ":50053"
    }

    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI not set")
    }

    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "orderdb"
    }

    inventoryURL := os.Getenv("INVENTORY_SERVICE_URL")
    if inventoryURL == "" {
        inventoryURL = "inventory-service:50052"
    }

    return &Config{
        MongoURI:           mongoURI,
        DatabaseName:       dbName,
        GRPCPort:          grpcPort,
        InventoryServiceURL: inventoryURL,
    }
}