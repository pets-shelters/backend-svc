package entity

import (
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type Shelter struct {
	ID          int64     `db:"id" structs:"-"`
	Name        string    `db:"name" structs:"name"`
	Logo        int64     `db:"logo" structs:"logo"`
	PhoneNumber string    `db:"phone_number" structs:"phone_number"`
	Instagram   *string   `db:"instagram" structs:"instagram"`
	Facebook    *string   `db:"facebook" structs:"facebook"`
	CreatedAt   date.Date `db:"created_at" structs:"created_at"`
}

type SheltersNames struct {
	ID   int64
	Name string
}

type UpdateShelter struct {
	Name        *string
	Logo        *int64
	PhoneNumber *string
	Instagram   *string
	Facebook    *string
}
