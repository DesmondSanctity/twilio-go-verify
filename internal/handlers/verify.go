package handlers

import (
	"encoding/json"
	"log"
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
	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.store.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := h.twilio.SendSMSOTP(user.PhoneNumber); err != nil {
		log.Printf("Failed to send SMS OTP: %v", err)
		http.Error(w, "Failed to send SMS OTP", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *VerifyHandler) VerifySMSOTP(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.store.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	verified, err := h.twilio.VerifySMSOTP(user.PhoneNumber, req.Code)
	if err != nil {
		log.Printf("Failed to verify SMS OTP: %v", err)
		http.Error(w, "Failed to verify SMS OTP", http.StatusInternalServerError)
		return
	}

	if !verified {
		http.Error(w, "Invalid OTP", http.StatusUnauthorized)
		return
	}

	user.SMSEnabled = true
	if err := h.store.UpdateUser(user); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *VerifyHandler) CreateTOTPFactor(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.store.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	sid, uri, err := h.twilio.CreateTOTPFactor(user.ID, user.Name)
	if err != nil {
		log.Printf("Failed to create TOTP factor: %v", err)
		http.Error(w, "Failed to create TOTP factor", http.StatusInternalServerError)
		return
	}

	user.TOTPFactorSid = sid
	if err := h.store.UpdateUser(user); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	response := QRResponse{
		QRCode:    uri,
		FactorSid: sid,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *VerifyHandler) VerifyFactor(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.store.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	verified, err := h.twilio.VerifyFactor(user.TOTPFactorSid, req.Code, user.ID)
	if err != nil {
		log.Printf("Failed to verify factor: %v", err)
		http.Error(w, "Failed to verify factor", http.StatusInternalServerError)
		return
	}

	if !verified {
		http.Error(w, "Invalid code", http.StatusUnauthorized)
		return
	}

	user.TOTPEnabled = true
	user.IsAuthenticated = true
	if err := h.store.UpdateUser(user); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *VerifyHandler) CreateTOTPChallenge(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.store.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	challengeSid, err := h.twilio.CreateTOTPChallenge(user.TOTPFactorSid, req.Code, user.ID)
	if err != nil {
		log.Printf("Failed to create TOTP challenge: %v", err)
		http.Error(w, "Failed to create TOTP challenge", http.StatusInternalServerError)
		return
	}

	user.IsAuthenticated = true
	if err := h.store.UpdateUser(user); err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"challengeSid": challengeSid})
}
