package requests

import "github.com/pets-shelters/backend-svc/pkg/date"

type TasksWithExecutionsFilters struct {
	Date *date.Date `form:"filter[date]"`
}
