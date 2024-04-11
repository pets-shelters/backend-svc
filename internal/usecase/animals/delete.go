package animals

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pkg/errors"
)

func (uc *UseCase) Delete(ctx context.Context, userId int64, animalId int64) error {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to get user entity")
	}
	if user == nil {
		return exceptions.NewPermissionDeniedException()
	}

	err = uc.repo.Transaction(ctx, func(tx pgx.Tx) error {
		locationId, err := uc.repo.GetAnimalsRepo().DeleteWithConn(ctx, tx, animalId)
		if err != nil {
			return errors.Wrap(err, "failed to delete animal by id")
		}
		if locationId == 0 {
			return exceptions.NewNotFoundException()
		}

		location, err := uc.repo.GetLocationsRepo().GetWithConn(ctx, tx, locationId)
		if err != nil {
			return errors.Wrap(err, "failed to get location entity")
		}
		if user.ShelterID.Int64 != location.ShelterID {
			return exceptions.NewPermissionDeniedException()
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to process transaction")
	}

	return nil
}
