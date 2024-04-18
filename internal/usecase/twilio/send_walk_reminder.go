package twilio

import (
	"encoding/json"
	"fmt"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
)

func (t *Twilio) SendWalkReminder(walkings []entity.WalkingReminder) error {
	for _, walking := range walkings {
		body := fmt.Sprintf("Нагадування від PetsShelters про заплановану прогулянку в притулку %s. На вас чекатиме %s %s о %s.",
			walking.ShelterName, walking.AnimalType, walking.AnimalName, walking.Time)
		params := &twilioApi.CreateMessageParams{}
		params.SetTo(walking.PhoneNumber)
		params.SetFrom(t.senderNumber)
		params.SetBody(body)

		resp, err := t.client.Api.CreateMessage(params)
		if err != nil {
			return errors.Wrap(err, "failed to send messages")
		}

		response, err := json.Marshal(*resp)
		log.Printf("response %+v", response)
		log.Printf("err %+v", err)
	}

	return nil
}
