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
	S3Bucket        string
}

func LoadConfig() *AppConfig {
	cfg := &AppConfig{
		MongoURI:        os.Getenv("MONGODB_URI"),
		MongoDatabase:   os.Getenv("MONGODB_DATABASE_NAME"),
		MongoCollection: os.Getenv("MONGODB_COLLECTION_NAME"),
		Env:             os.Getenv("ENV"),
		S3Bucket:        "my-cross-account-bucket-10032025",
	}

	if cfg.MongoURI == "" || cfg.MongoDatabase == "" || cfg.MongoCollection == "" {
		log.Fatal("Missing required MongoDB environment variables")
	}

	if cfg.S3Bucket == "" {
		log.Fatal("Missing required environment variable S3_BUCKET_NAME")
	}

	return cfg
}
