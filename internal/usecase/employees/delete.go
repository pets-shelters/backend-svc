package employees

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
		userDeleted, err := uc.repo.GetUsersRepo().DeleteWithConn(ctx, tx, idToDelete)
		if err != nil {
			return errors.Wrap(err, "failed to delete user entity")
		}
		if userDeleted == nil {
			return exceptions.NewNotFoundException()
		}
		if userDeleted.ShelterID != user.ShelterID {
			return exceptions.NewPermissionDeniedException()
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to process transaction")
	}

	return nil
}
