package requests

import (
	customTime "github.com/pets-shelters/backend-svc/pkg/custom_time"
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type ApproveWalking struct {
	Date *date.Date          `json:"date,omitempty"`
	Time customTime.NullTime `json:"time" validate:"required"`
}
