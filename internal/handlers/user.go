package handlers

import (
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

}
