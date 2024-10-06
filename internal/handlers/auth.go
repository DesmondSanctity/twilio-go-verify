package handlers

import (
	"net/http"

	"github.com/desmomndsanctity/twilio-go-verify/internal/store"
	"github.com/desmomndsanctity/twilio-go-verify/internal/twilio"
)

type AuthHandler struct {
	store  *store.InMemoryStore
	twilio *twilio.TwilioVerify
}

func NewAuthHandler(store *store.InMemoryStore, twilio *twilio.TwilioVerify) *AuthHandler {
	return &AuthHandler{
		store:  store,
		twilio: twilio,
	}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {

}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {

}
