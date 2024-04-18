package walkings

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) CreateApproved(ctx context.Context, req requests.CreateApprovedWalking, animalId int64, userId int64) error {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to get user entity")
	}
	if user == nil {
		return exceptions.NewPermissionDeniedException()
	}

	shelterId, err := uc.repo.GetAnimalsRepo().SelectShelterID(ctx, animalId)
	if err != nil {
		return errors.Wrap(err, "failed to get animal entity")
	}
	if shelterId == 0 {
		return exceptions.NewNotFoundException()
	}
	if shelterId != user.ShelterID.Int64 {
		return exceptions.NewPermissionDeniedException()
	}

	_, err = uc.repo.GetWalkingsRepo().Create(ctx, entity.Walking{
		Status:      structs.ApprovedWalkingStatus,
		AnimalID:    animalId,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Date:        req.Date,
		Time:        req.Time,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create walking entity")
	}

	return nil
}
