package responses

import (
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type Shelter struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Logo        string    `json:"logo"`
	PhoneNumber string    `json:"phone_number"`
	Instagram   *string   `json:"instagram,omitempty"`
	Facebook    *string   `json:"facebook,omitempty"`
	CreatedAt   date.Date `json:"created_at"`
}
