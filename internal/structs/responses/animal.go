package responses

import (
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type Animal struct {
	ID        int64     `json:"id"`
	Photo     string    `json:"photo"`
	Name      string    `json:"name"`
	BirthDate date.Date `json:"birth_date"`
	Type      string    `json:"type"`
}
