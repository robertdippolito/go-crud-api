package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"k8s-api/config"
	"k8s-api/db"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Handler struct {
	Config *config.AppConfig
}

type AppConfig struct {
	MongoDatabase   string
	MongoCollection string
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	collection := db.Client.Database(h.Config.MongoDatabase).Collection(h.Config.MongoCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var users []User
	if err := cursor.All(ctx, &users); err != nil {
		http.Error(w, "Error decoding user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	collection := db.Client.Database(h.Config.MongoDatabase).Collection(h.Config.MongoCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
