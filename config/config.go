package config

import (
	"log"
	"os"
)

type AppConfig struct {
	MongoURI        string
	MongoDatabase   string
	MongoCollection string
	Env             string
}

func LoadConfig() *AppConfig {
	cfg := &AppConfig{
		MongoURI:        os.Getenv("MONGODB_URI"),
		MongoDatabase:   os.Getenv("MONGODB_DATABASE_NAME"),
		MongoCollection: os.Getenv("MONGODB_COLLECTION_NAME"),
		Env:             os.Getenv("ENV"),
	}

	log.Println("Config values:")
	log.Printf("  MONGODB_URI: %t\n", cfg.MongoURI != "")
	log.Printf("  MONGODB_DATABASE_NAME: %t\n", cfg.MongoDatabase != "")
	log.Printf("  MONGODB_COLLECTION_NAME: %t\n", cfg.MongoCollection != "")
	log.Printf("  ENV: %s\n", cfg.Env)

	if cfg.MongoURI == "" || cfg.MongoDatabase == "" || cfg.MongoCollection == "" {
		log.Fatal("Missing required environment variables")
	}

	return cfg
}
