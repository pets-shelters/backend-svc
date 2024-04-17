package requests

import (
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type AnimalsFilters struct {
	ShelterID     []int64    `form:"filter[shelter_id]"`
	LocationID    []int64    `form:"filter[location_id]"`
	Gender        *string    `form:"filter[gender]" validate:"omitempty,oneof=female male"`
	Adopted       *bool      `form:"filter[adopted]"`
	Sterilized    *bool      `form:"filter[sterilized]"`
	ForAdoption   *bool      `form:"filter[for_adoption]"`
	ForWalking    *bool      `form:"filter[for_walking]"`
	BirthDateFrom *date.Date `form:"filter[birth_date_from]"`
	BirthDateTo   *date.Date `form:"filter[birth_date_to]"`
	Type          []string   `form:"filter[type]"`
	Name          *string    `form:"filter[name]"`
	City          []string   `form:"filter[city]"`
}
