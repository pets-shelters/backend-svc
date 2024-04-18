package entity

import (
	"github.com/pets-shelters/backend-svc/internal/structs"
	customTime "github.com/pets-shelters/backend-svc/pkg/custom_time"
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type Walking struct {
	ID          int64                 `db:"id" structs:"-"`
	Status      structs.WalkingStatus `db:"status" structs:"status"`
	AnimalID    int64                 `db:"animal_id" structs:"animal_id"`
	Name        string                `db:"name" structs:"name"`
	PhoneNumber string                `db:"phone_number" structs:"phone_number"`
	Date        date.Date             `db:"date" structs:"date"`
	Time        customTime.NullTime   `db:"time" structs:"time"`
}

type WalkingsFilters struct {
	Status    *structs.WalkingStatus
	AnimalID  *int64
	ShelterID *int64
	Date      *date.Date
}

type WalkingReminder struct {
	AnimalName  string
	AnimalType  string
	PhoneNumber string
	ShelterName string
	Time        customTime.NullTime
}

type UpdateWalking struct {
	Status *structs.WalkingStatus
	Date   *date.Date
	Time   *customTime.NullTime
}
