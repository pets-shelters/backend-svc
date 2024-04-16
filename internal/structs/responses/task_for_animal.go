package responses

import (
	customTime "github.com/pets-shelters/backend-svc/pkg/custom_time"
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type TaskForAnimal struct {
	ID               int64               `json:"id"`
	Description      string              `json:"description"`
	StartDate        date.Date           `json:"start_date"`
	EndDate          date.Date           `json:"end_date"`
	Time             customTime.NullTime `json:"time,omitempty"`
	ExecutionsNumber int64               `json:"executions_number"`
}
