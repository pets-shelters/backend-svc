package entity

import (
	customTime "github.com/pets-shelters/backend-svc/pkg/custom_time"
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type Task struct {
	ID          int64            `db:"id" structs:"-"`
	Description string           `db:"description" structs:"description"`
	StartDate   date.Date        `db:"start_date" structs:"start_date"`
	EndDate     date.Date        `db:"end_date" structs:"end_date"`
	Time        *customTime.Time `db:"time" structs:"time"`
}
