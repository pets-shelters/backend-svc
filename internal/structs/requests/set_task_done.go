package requests

import "github.com/pets-shelters/backend-svc/pkg/date"

type SetTaskDone struct {
	Date date.Date `json:"date" validate:"required"`
}
