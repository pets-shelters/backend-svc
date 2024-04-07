package entity

import (
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type Animal struct {
	ID                 int64                `db:"id" structs:"-"`
	LocationID         int64                `db:"location_id" structs:"location_id"`
	Photo              int64                `db:"photo" structs:"photo"`
	Name               string               `db:"name" structs:"name"`
	BirthDate          date.Date            `db:"birth_date" structs:"birth_date"`
	Type               string               `db:"type" structs:"type"`
	Gender             structs.AnimalGender `db:"gender" structs:"gender"`
	Sterilized         bool                 `db:"sterilized" structs:"sterilized"`
	PrivateDescription *string              `db:"private_description" structs:"private_description"`
	PublicDescription  *string              `db:"public_description" structs:"public_description"`
}
