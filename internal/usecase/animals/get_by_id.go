package animals

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetById(ctx context.Context, animalId int64, userId *int64) (*responses.Animal, error) {
	animal, err := uc.repo.GetAnimalsRepo().Get(ctx, animalId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get animal entity")
	}
	if animal == nil {
		return nil, exceptions.NewNotFoundException()
	}

	file, err := uc.repo.GetFilesRepo().Get(ctx, animal.Photo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get file entity")
	}

	response := responses.Animal{
		ID:                animal.ID,
		LocationID:        animal.LocationID,
		Photo:             uc.s3Endpoint + "/" + file.Bucket + file.Path,
		Name:              animal.Name,
		BirthDate:         animal.BirthDate,
		Type:              animal.Type,
		Gender:            animal.Gender,
		Sterilized:        animal.Sterilized,
		ForAdoption:       animal.ForAdoption,
		ForWalking:        animal.ForWalking,
		PublicDescription: animal.PublicDescription,
	}
	if userId != nil {
		user, err := uc.repo.GetUsersRepo().Get(ctx, *userId)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get user entity")
		}
		location, err := uc.repo.GetLocationsRepo().Get(ctx, animal.LocationID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get location entity")
		}
		if user.ShelterID.Int64 == location.ShelterID {
			response.PrivateDescription = animal.PrivateDescription
			if animal.AdopterID.Valid {
				response.AdopterID = &animal.AdopterID.Int64
			}
		}
	}

	return &response, nil
}
