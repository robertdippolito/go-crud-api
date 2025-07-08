package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"time"
)

type Calculation struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Result struct {
	Sum float32 `json:"sum"`
}

func (h *Handler) Compute(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) BurnTest(w http.ResponseWriter, r *http.Request) {
	end := time.Now().Add(1 * time.Second)
	for time.Now().Before(end) {
		math.Sqrt(12345.6789)
	}
	w.Write([]byte("done"))
}
