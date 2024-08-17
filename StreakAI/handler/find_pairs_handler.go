package handler

import (
	"Task/service"
	"Task/types"
	"encoding/json"
	"errors"
	"net/http"
)

func validateInput(req types.Request) error {
	if req.Numbers == nil || len(req.Numbers) == 0 {
		return http.ErrBodyNotAllowed
	}

	for _, num := range req.Numbers {
		if num != int(num) {
			return errors.New("invalid Integer")
		}
	}
	return nil

}

func FindPairsHandler(w http.ResponseWriter, r *http.Request) {
	var req types.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := validateInput(req); err != nil {
		http.Error(w, "Invalid Input :"+err.Error(), http.StatusBadRequest)
		return
	}
	solutions := service.FindPairs(req.Numbers, req.Target)

	resp := types.Response{Solutions: solutions}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
