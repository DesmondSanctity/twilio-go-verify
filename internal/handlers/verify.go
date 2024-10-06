package handlers

import (
	"net/http"

	"github.com/desmomndsanctity/twilio-go-verify/internal/store"
	"github.com/desmomndsanctity/twilio-go-verify/internal/twilio"
)

type VerifyHandler struct {
	store  *store.InMemoryStore
	twilio *twilio.TwilioVerify
}

type QRResponse struct {
	QRCode    string `json:"qrCode"`
	FactorSid string `json:"factorSid"`
}

func NewVerifyHandler(store *store.InMemoryStore, twilio *twilio.TwilioVerify) *VerifyHandler {
	return &VerifyHandler{
		store:  store,
		twilio: twilio,
	}
}

func (h *VerifyHandler) SendSMSOTP(w http.ResponseWriter, r *http.Request) {

}

func (h *VerifyHandler) VerifySMSOTP(w http.ResponseWriter, r *http.Request) {

}

func (h *VerifyHandler) CreateTOTPFactor(w http.ResponseWriter, r *http.Request) {

}

func (h *VerifyHandler) VerifyFactor(w http.ResponseWriter, r *http.Request) {

}

func (h *VerifyHandler) CreateTOTPChallenge(w http.ResponseWriter, r *http.Request) {

}
