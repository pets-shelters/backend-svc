package responses

import (
	customTime "github.com/pets-shelters/backend-svc/pkg/custom_time"
	"github.com/pets-shelters/backend-svc/pkg/date"
	"time"
)

type TaskWithExecutions struct {
	ID          int64               `json:"id"`
	Description string              `json:"description"`
	StartDate   date.Date           `json:"start_date"`
	EndDate     date.Date           `json:"end_date"`
	Time        customTime.NullTime `json:"time,omitempty"`
	Executions  []TaskExecution     `json:"executions"`
}

type TaskExecution struct {
	UserID *int64    `json:"user_id,omitempty"`
	Date   date.Date `json:"date"`
	DoneAt time.Time `json:"done_at"`
}
