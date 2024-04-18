package requests

import (
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/pkg/date"
)

type WalkingsFilters struct {
	Status   *structs.WalkingStatus `form:"filter[status]" validate:"omitempty,oneof=pending approved"`
	AnimalID *int64                 `form:"filter[animal_id]"`
	Date     *date.Date             `form:"filter[date]"`
}
