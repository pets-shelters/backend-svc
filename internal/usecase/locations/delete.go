package locations

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pkg/errors"
)

func (uc *UseCase) Delete(ctx context.Context, userId int64, idToDelete int64) error {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to get user entity")
	}
	if user.Role != structs.ManagerUserRole {
		return exceptions.NewPermissionDeniedException()
	}

	err = uc.repo.Transaction(ctx, func(tx pgx.Tx) error {
		locationDeleted, err := uc.repo.GetLocationsRepo().DeleteWithConn(ctx, tx, idToDelete)
		if err != nil {
			return errors.Wrap(err, "failed to delete location entity")
		}
		if locationDeleted == nil {
			return exceptions.NewNotFoundException()
		}
		if locationDeleted.ShelterID != user.ShelterID.Int64 {
			return exceptions.NewPermissionDeniedException()
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to process transaction")
	}

	return nil
}
