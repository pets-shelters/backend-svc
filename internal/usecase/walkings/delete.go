package walkings

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pkg/errors"
)

func (uc *UseCase) Delete(ctx context.Context, userId int64, walkingId int64) error {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to get user entity")
	}
	if user == nil {
		return exceptions.NewPermissionDeniedException()
	}

	err = uc.repo.Transaction(ctx, func(tx pgx.Tx) error {
		walking, err := uc.repo.GetWalkingsRepo().DeleteWithConn(ctx, tx, walkingId)
		if err != nil {
			return errors.Wrap(err, "failed to delete animal by id")
		}
		if walking == nil {
			return exceptions.NewNotFoundException()
		}

		shelterId, err := uc.repo.GetAnimalsRepo().SelectShelterIDWithConn(ctx, tx, walking.AnimalID)
		if err != nil {
			return errors.Wrap(err, "failed to get animal entity")
		}
		if shelterId != user.ShelterID.Int64 {
			return exceptions.NewPermissionDeniedException()
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to process transaction")
	}

	return nil
}
