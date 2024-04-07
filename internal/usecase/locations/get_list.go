package locations

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetList(ctx context.Context, shelterId int64) ([]responses.Location, error) {
	locationsAnimals, err := uc.repo.GetLocationsRepo().SelectWithAnimals(ctx, shelterId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select locationsAnimals entities")
	}

	return formLocationsResponse(locationsAnimals), nil
}

func formLocationsResponse(locationsAnimals []entity.LocationsAnimals) []responses.Location {
	response := make([]responses.Location, 0)
	for _, locationAnimal := range locationsAnimals {
		response = append(response, responses.Location{
			ID:            locationAnimal.ID,
			City:          locationAnimal.City,
			Address:       locationAnimal.Address,
			ShelterID:     locationAnimal.ShelterID,
			AnimalsNumber: locationAnimal.AnimalsNumber,
		})
	}

	return response
}
