package handlers

import (
	"encoding/json"
	"net/http"
)

type Calculation struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Result struct {
	Sum float32 `json:"sum"`
}

func Compute(w http.ResponseWriter, r *http.Request) {
	var calc Calculation
	err := json.NewDecoder(r.Body).Decode(&calc)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	result := Result{Sum: calc.X + calc.Y}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
