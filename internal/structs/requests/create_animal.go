package requests

import (
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type CreateAnimal struct {
	LocationID         int64     `json:"location_id" validate:"required"`
	Photo              int64     `json:"photo" validate:"required"`
	Name               string    `json:"name" validate:"required"`
	BirthDate          date.Date `json:"birth_date" validate:"required"`
	Type               string    `json:"type" validate:"required"`
	Gender             string    `json:"gender" validate:"required,oneof=female male"`
	Sterilized         bool      `json:"sterilized" validate:"required"`
	ForAdoption        bool      `json:"for_adoption" validate:"required"`
	ForWalking         bool      `json:"for_walking" validate:"required"`
	PrivateDescription *string   `json:"private_description" validate:"omitempty"`
	PublicDescription  *string   `json:"public_description" validate:"omitempty"`
}
