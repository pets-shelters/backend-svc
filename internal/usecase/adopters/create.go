package adopters

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) Create(ctx context.Context, req requests.CreateAdopter, animalId int64) (int64, error) {
	var adopterId int64
	var err error
	adopterId, err = uc.repo.GetAdoptersRepo().Create(ctx, entity.Adopter{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		if !errors.As(err, &exceptions.AdopterExistsException{}) {
			return 0, errors.Wrap(err, "failed to create adopter entity")
		}

		adopter, err := uc.repo.GetAdoptersRepo().GetByPhoneNumber(ctx, req.PhoneNumber)
		if err != nil {
			return 0, errors.Wrap(err, "failed to get adopter entity")
		}
		adopterId = adopter.ID
	}

	rowsAffected, err := uc.repo.GetAnimalsRepo().Update(ctx, animalId, entity.UpdateAnimal{
		AdopterID: &adopterId,
	})
	if err != nil {
		return 0, errors.Wrap(err, "failed to set adopter for animal")
	}
	if rowsAffected == 0 {
		return 0, exceptions.NewNotFoundException()
	}

	return adopterId, nil
}
