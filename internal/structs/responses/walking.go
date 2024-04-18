package responses

import (
	"github.com/pets-shelters/backend-svc/internal/structs"
	customTime "github.com/pets-shelters/backend-svc/pkg/custom_time"
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type Walking struct {
	ID          int64                 `json:"id"`
	Status      structs.WalkingStatus `json:"status"`
	AnimalID    int64                 `json:"animal_id"`
	Name        string                `json:"name"`
	PhoneNumber string                `json:"phone_number"`
	Date        date.Date             `json:"date"`
	Time        customTime.NullTime   `json:"time"`
}
