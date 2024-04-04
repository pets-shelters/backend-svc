package mailjet

import (
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/pets-shelters/backend-svc/configs"
)

type Mailjet struct {
	client *mailjet.Client
	cfg    configs.Mailjet
}

func NewMailjet(cfg configs.Mailjet) *Mailjet {
	client := mailjet.NewMailjetClient(cfg.PublicKey, cfg.PrivateKey)
	return &Mailjet{
		client,
		cfg,
	}
}
