package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/desmomndsanctity/twilio-go-verify/internal/store"
)

type UserHandler struct {
	store *store.InMemoryStore
}

func NewUserHandler(store *store.InMemoryStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	user, err := h.store.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// if !user.IsAuthenticated {
	// 	http.Error(w, "User is not authenticated", http.StatusUnauthorized)
	// 	return
	// }

	response := struct {
		Name            string `json:"name"`
		Email           string `json:"email"`
		SMSEnabled      bool   `json:"smsEnabled"`
		TOTPEnabled     bool   `json:"totpEnabled"`
		TOTPFactorSid   string `json:"totpFactorSid"`
		IsAuthenticated bool   `json:"isAuthenticated"`
	}{
		Name:            user.Name,
		Email:           user.Email,
		SMSEnabled:      user.SMSEnabled,
		TOTPEnabled:     user.TOTPEnabled,
		TOTPFactorSid:   user.TOTPFactorSid,
		IsAuthenticated: user.IsAuthenticated,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
