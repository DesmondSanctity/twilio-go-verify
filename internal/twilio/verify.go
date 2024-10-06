package twilio

import (
	"github.com/twilio/twilio-go"
)

type TwilioVerify struct {
	client *twilio.RestClient
	sid    string
}

func NewTwilioVerify(accountSid, authToken, verifySid string) *TwilioVerify {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	return &TwilioVerify{
		client: client,
		sid:    verifySid,
	}
}

func (tv *TwilioVerify) SendSMSOTP(to string) error {
	return nil
}

func (tv *TwilioVerify) VerifySMSOTP(to, code string) (bool, error) {
	return false, nil
}

func (tv *TwilioVerify) CreateTOTPFactor(identity, name string) (string, string, error) {
	return "", "", nil
}

func (tv *TwilioVerify) VerifyFactor(factorSid, code string, identity string) (bool, error) {
	return false, nil
}

func (tv *TwilioVerify) CreateTOTPChallenge(factorSid string, code string, identity string) (string, error) {
	return "", nil
}
