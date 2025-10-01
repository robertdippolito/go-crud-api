package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"k8s-api/config"
	"k8s-api/db"
	"k8s-api/handlers"

	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	env := os.Getenv("ENV")
	if env == "" {
		env = "production" // default fallback
	}

	log.Printf("Running in %s environment\n", env)

	conf := config.LoadConfig()

	awsCfg, err := awscfg.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal("Failed to load AWS configuration", err)
	}

	handler := &handlers.Handler{
		Config:   conf,
		S3Client: s3.NewFromConfig(awsCfg),
	}

	err = db.InitMongoDB(conf.MongoURI)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}

	r := NewRouter(handler)
	log.Println("Server running on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
