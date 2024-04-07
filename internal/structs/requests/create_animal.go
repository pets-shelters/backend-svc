package requests

import (
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type CreateAnimal struct {
	LocationID         int64                `json:"location_id" validate:"required"`
	Photo              int64                `json:"photo" validate:"required"`
	Name               string               `json:"name" validate:"required"`
	BirthDate          date.Date            `json:"birth_date" validate:"required"`
	Type               string               `json:"type" validate:"required"`
	Gender             structs.AnimalGender `json:"gender" validate:"required"`
	Sterilized         bool                 `json:"sterilized" validate:"required"`
	PrivateDescription *string              `json:"private_description" validate:"omitempty"`
	PublicDescription  *string              `json:"public_description" validate:"omitempty"`
}
