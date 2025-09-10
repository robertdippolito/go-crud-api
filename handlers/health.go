package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func (h *Handler) GetHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{Message: "I am healthy! HELLO YOUTUBE"})
}
