package requests

import "github.com/pets-shelters/backend-svc/pkg/date"

type UpdateAnimal struct {
	LocationID         *int64     `json:"location_id"`
	Photo              *int64     `json:"photo"`
	Name               *string    `json:"name"`
	BirthDate          *date.Date `json:"birth_date"`
	Type               *string    `json:"type"`
	Gender             *string    `json:"gender" validate:"omitempty,oneof=female male"`
	Sterilized         *bool      `json:"sterilized"`
	ForAdoption        *bool      `json:"for_adoption"`
	ForWalking         *bool      `json:"for_walking"`
	AdopterID          *int64     `json:"adopter_id"`
	PublicDescription  *string    `json:"public_description"`
	PrivateDescription *string    `json:"private_description"`
}
