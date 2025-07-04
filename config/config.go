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
		MongoURI:        os.Getenv("MONGO_URI"),
		MongoDatabase:   os.Getenv("MONGO_DATABASE_NAME"),
		MongoCollection: os.Getenv("MONGO_COLLECTION_NAME"),
		Env:             os.Getenv("ENV"),
	}

	if cfg.MongoURI == "" || cfg.MongoDatabase == "" || cfg.MongoCollection == "" {
		log.Fatal("Missing required environment variables")
	}

	return cfg
}
