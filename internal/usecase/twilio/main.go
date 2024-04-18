package twilio

import (
	"github.com/pets-shelters/backend-svc/configs"
	"github.com/twilio/twilio-go"
)

type Twilio struct {
	client       *twilio.RestClient
	senderNumber string
}

func NewTwilio(cfg configs.Twilio) *Twilio {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.AccountSid,
		Password: cfg.AuthToken,
	})
	return &Twilio{
		client:       client,
		senderNumber: cfg.SenderNumber,
	}
}
