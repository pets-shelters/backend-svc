package requests

import (
	customTime "github.com/pets-shelters/backend-svc/pkg/custom_time"
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type CreateTask struct {
	AnimalID    int64            `json:"animal_id" validate:"required"`
	Description string           `json:"description" validate:"required"`
	StartDate   date.Date        `json:"start_date" validate:"required"`
	EndDate     date.Date        `json:"end_date" validate:"required"`
	Time        *customTime.Time `json:"time"`
}
