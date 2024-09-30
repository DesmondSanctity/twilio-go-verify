package twilio

import (
	"fmt"

	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
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
	params := &verify.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	_, err := tv.client.VerifyV2.CreateVerification(tv.sid, params)
	return err
}

func (tv *TwilioVerify) VerifySMSOTP(to, code string) (bool, error) {
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)

	resp, err := tv.client.VerifyV2.CreateVerificationCheck(tv.sid, params)
	if err != nil {
		return false, err
	}

	return *resp.Status == "approved", nil
}

func (tv *TwilioVerify) CreateTOTPFactor(identity, name string) (string, string, error) {
	params := &verify.CreateNewFactorParams{}
	params.SetFriendlyName(name + "'s totp")
	params.SetFactorType("totp")

	resp, err := tv.client.VerifyV2.CreateNewFactor(tv.sid, identity, params)
	if err != nil {
		return "", "", err
	}

	binding, ok := (*resp.Binding).(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("unexpected binding type")
	}

	uri, ok := binding["uri"].(string)
	if !ok {
		return "", "", fmt.Errorf("uri not found in binding or not a string")
	}

	return *resp.Sid, uri, nil
}

func (tv *TwilioVerify) VerifyFactor(factorSid, code string, identity string) (bool, error) {
	params := &verify.UpdateFactorParams{}
	params.SetAuthPayload(code)

	resp, err := tv.client.VerifyV2.UpdateFactor(tv.sid, identity, factorSid, params)
	if err != nil {
		return false, err
	}

	return *resp.Status == "verified", nil
}

func (tv *TwilioVerify) CreateTOTPChallenge(factorSid string, code string, identity string) (string, error) {
	params := &verify.CreateChallengeParams{}
	params.SetAuthPayload(code)
	params.SetFactorSid(factorSid)

	resp, err := tv.client.VerifyV2.CreateChallenge(tv.sid, identity, params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}
