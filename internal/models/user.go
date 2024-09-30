package models

type User struct {
	ID              string
	Name            string
	Email           string
	Password        string
	PhoneNumber     string
	IsAuthenticated bool
	SMSEnabled      bool
	TOTPFactorSid   string
	TOTPEnabled     bool
}
