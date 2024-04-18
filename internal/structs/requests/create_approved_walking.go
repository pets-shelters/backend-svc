package requests

import (
	customTime "github.com/pets-shelters/backend-svc/pkg/custom_time"
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type CreateApprovedWalking struct {
	Name        string              `json:"name" validate:"required"`
	PhoneNumber string              `json:"phone_number" validate:"required,len=13"`
	Date        date.Date           `json:"date" validate:"required"`
	Time        customTime.NullTime `json:"time" validate:"required"`
}
