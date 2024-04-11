package responses

import (
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type AnimalForList struct {
	ID        int64     `json:"id"`
	Photo     string    `json:"photo"`
	Name      string    `json:"name"`
	BirthDate date.Date `json:"birth_date"`
	Type      string    `json:"type"`
}

type Animal struct {
	ID                 int64     `json:"id"`
	LocationID         int64     `json:"location_id"`
	Photo              string    `json:"photo"`
	Name               string    `json:"name"`
	BirthDate          date.Date `json:"birth_date"`
	Type               string    `json:"type"`
	Gender             string    `json:"gender"`
	Sterilized         bool      `json:"sterilized"`
	AdopterID          *int64    `json:"adopter_id,omitempty"`
	PublicDescription  *string   `json:"public_description,omitempty"`
	PrivateDescription *string   `json:"private_description,omitempty"`
}
