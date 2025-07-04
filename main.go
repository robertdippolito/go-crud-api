package main

import (
	"log"
	"net/http"
	"os"

	"k8s-api/config"
	"k8s-api/db"
	"k8s-api/handlers"

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

	handler := &handlers.Handler{Config: conf}

	err := db.InitMongoDB(conf.MongoURI)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}

	r := NewRouter(handler)
	log.Println("Server running on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
