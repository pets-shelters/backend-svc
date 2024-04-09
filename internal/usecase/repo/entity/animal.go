package entity

import (
	"github.com/pets-shelters/backend-svc/pkg/date"
	"time"
)

type Animal struct {
	ID                 int64     `db:"id" structs:"-"`
	LocationID         int64     `db:"location_id" structs:"location_id"`
	Photo              int64     `db:"photo" structs:"photo"`
	Name               string    `db:"name" structs:"name"`
	BirthDate          date.Date `db:"birth_date" structs:"birth_date"`
	Type               string    `db:"type" structs:"type"`
	Gender             string    `db:"gender" structs:"gender"`
	Sterilized         bool      `db:"sterilized" structs:"sterilized"`
	PrivateDescription *string   `db:"private_description" structs:"private_description"`
	PublicDescription  *string   `db:"public_description" structs:"public_description"`
}

type AnimalForList struct {
	ID          int64
	PhotoBucket string
	PhotoPath   string
	Name        string
	BirthDate   time.Time
	Type        string
}

type AnimalsFilters struct {
	ShelterID     []int64
	LocationID    []int64
	Gender        *string
	Sterilized    *bool
	BirthDateFrom *date.Date
	BirthDateTo   *date.Date
	Type          []string
	Name          *string
	City          []string
}
