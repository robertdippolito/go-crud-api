package main

import (
	"log"
	"net/http"
	"os"

	"k8s-api/db"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	err = db.InitMongoDB(mongoURI)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}

	r := NewRouter()
	log.Println("Server running on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
