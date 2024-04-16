package mailjet

import (
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/pkg/date"
	"github.com/pkg/errors"
)

func (m *Mailjet) SendTasksEmail(toEmail string, date date.Date, tasks []structs.TaskForEmail) error {
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
			TemplateID:       m.cfg.TasksTemplateId,
			TemplateLanguage: true,
			Variables: map[string]interface{}{
				"date":  date.String(),
				"url":   m.cfg.TasksUrl,
				"items": tasks,
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
