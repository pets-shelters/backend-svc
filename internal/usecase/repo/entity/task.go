package entity

import (
	customTime "github.com/pets-shelters/backend-svc/pkg/custom_time"
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type Task struct {
	ID          int64               `db:"id" structs:"-"`
	Description string              `db:"description" structs:"description"`
	StartDate   date.Date           `db:"start_date" structs:"start_date"`
	EndDate     date.Date           `db:"end_date" structs:"end_date"`
	Time        customTime.NullTime `db:"time" structs:"time"`
}

type TasksFilters struct {
	Date      *date.Date
	ShelterID *int64
}

type TaskWithExecutions struct {
	Task
	Executions []TaskExecutionForList
}

type TaskForAnimal struct {
	ID               int64
	Description      string
	StartDate        date.Date
	EndDate          date.Date
	Time             customTime.NullTime
	ExecutionsNumber int64
}

type EmployeeTasks struct {
	EmployeeEmail string
	Tasks         []TaskForEmail
}

type TaskForEmail struct {
	Description string
	AnimalType  string
	AnimalName  string
	Time        customTime.NullTime
}
