package handlers

import (
	"encoding/json"
	"net/http"
)

type AddRequest struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

type AddResponse struct {
	Result int `json:"result"`
}

func Add(w http.ResponseWriter, r *http.Request) {
	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := req.Number1 + req.Number2
	resp := AddResponse{Result: result}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
