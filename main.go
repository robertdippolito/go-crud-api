package main

import (
	"log"
	"net/http"
	"os"

	"k8s-api/config"
	"k8s-api/db"
	"k8s-api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	mongoDB := os.Getenv("MONGO_DATABASE_NAME")
	mongoCollection := os.Getenv("MONGO_COLLECTION_NAME")

	if mongoURI == "" || mongoDB == "" || mongoCollection == "" {
		log.Fatal("Required environment variables are missing")
	}

	err := db.InitMongoDB(mongoURI)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}

	conf := &config.AppConfig{
		MongoDatabase:   mongoDB,
		MongoCollection: mongoCollection,
	}

	h := &handlers.Handler{Config: conf}

	router := mux.NewRouter()
	router.HandleFunc("/users", h.GetUsers).Methods("GET")
	router.HandleFunc("/users", h.CreateUser).Methods("POST")

	log.Println("Server running on :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
