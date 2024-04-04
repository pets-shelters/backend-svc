package mailjet

import (
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/pkg/errors"
)

func (m *Mailjet) SendInvitationEmail(shelterName string, toEmail string) error {
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: m.cfg.SenderEmail,
				Name:  m.cfg.SenderName,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: toEmail,
				},
			},
			TemplateID:       5842009,
			TemplateLanguage: true,
			Variables: map[string]interface{}{
				"shelter_name": shelterName,
				"login_url":    m.cfg.InvitationUrl,
			},
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := m.client.SendMailV31(&messages)
	if err != nil {
		return errors.Wrap(err, "failed to send email")
	}

	return nil
}
