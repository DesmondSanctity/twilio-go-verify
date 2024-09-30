package main

import (
	"log"
	"net/http"
	"os"

	"github.com/desmomndsanctity/twilio-go-verify/internal/handlers"
	"github.com/desmomndsanctity/twilio-go-verify/internal/store"
	"github.com/desmomndsanctity/twilio-go-verify/internal/twilio"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the in-memory store
	inMemoryStore := store.NewInMemoryStore()

	// Initialize Twilio Verify client
	twilioVerify := twilio.NewTwilioVerify(
		os.Getenv("TWILIO_ACCOUNT_SID"),
		os.Getenv("TWILIO_AUTH_TOKEN"),
		os.Getenv("TWILIO_VERIFY_SID"),
	)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(inMemoryStore, twilioVerify)
	verifyHandler := handlers.NewVerifyHandler(inMemoryStore, twilioVerify)
	userHandler := handlers.NewUserHandler(inMemoryStore)

	// Set up router
	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/api/signup", authHandler.SignUp).Methods("POST")
	r.HandleFunc("/api/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/api/logout", authHandler.Logout).Methods("POST")
	r.HandleFunc("/api/verify/send-sms", verifyHandler.SendSMSOTP).Methods("POST")
	r.HandleFunc("/api/verify/verify-sms", verifyHandler.VerifySMSOTP).Methods("POST")
	r.HandleFunc("/api/verify/create-totp", verifyHandler.CreateTOTPFactor).Methods("POST")
	r.HandleFunc("/api/verify/verify-factor", verifyHandler.VerifyFactor).Methods("POST")
	r.HandleFunc("/api/verify/create-totp-challenge", verifyHandler.CreateTOTPChallenge).Methods("POST")

	// User routes
	r.HandleFunc("/api/user", userHandler.GetUserInfo).Methods("GET")

	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	// Start the server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
