package main

import (
	"log"
	"net/http"
	"os"

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

	// Set up router
	r := mux.NewRouter()

	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	// Start the server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
