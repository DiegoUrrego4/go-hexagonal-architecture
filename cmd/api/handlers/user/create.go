package user

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time-tracker/internal/user/domain"
)

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest domain.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusInternalServerError)
		return
	}

	user := domain.User{
		ID:       uuid.New().String(),
		Name:     userRequest.Name,
		LastName: userRequest.LastName,
		Age:      userRequest.Age,
	}

	err = h.UserService.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": user.ID}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
