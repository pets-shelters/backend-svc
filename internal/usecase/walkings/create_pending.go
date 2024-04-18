package walkings

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) CreatePending(ctx context.Context, req requests.CreatePendingWalking, animalId int64) error {
	animal, err := uc.repo.GetAnimalsRepo().Get(ctx, animalId)
	if err != nil {
		return errors.Wrap(err, "failed to get animal entity")
	}
	if animal == nil {
		return exceptions.NewNotFoundException()
	}
	if !animal.ForWalking {
		return exceptions.NewAnimalUnavailableForWalkingException()
	}

	_, err = uc.repo.GetWalkingsRepo().Create(ctx, entity.Walking{
		Status:      structs.PendingWalkingStatus,
		AnimalID:    animalId,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Date:        req.Date,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create walking entity")
	}

	return nil
}
