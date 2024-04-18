package requests

import (
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type CreatePendingWalking struct {
	Name        string    `json:"name" validate:"required"`
	PhoneNumber string    `json:"phone_number" validate:"required,len=13"`
	Date        date.Date `json:"date" validate:"required"`
}
