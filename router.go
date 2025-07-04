package main

import (
	"k8s-api/handlers"

	"github.com/gorilla/mux"
)

func NewRouter(h *handlers.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", h.GetHealth).Methods("GET")
	r.HandleFunc("/users", h.GetUsers).Methods("GET")
	r.HandleFunc("/users", h.CreateUser).Methods("POST")
	r.HandleFunc("/compute", h.Compute).Methods("POST")

	return r
}
