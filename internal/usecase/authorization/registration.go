package authorization

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) Registration(ctx context.Context, shelter entity.Shelter, email string) error {
	err := uc.repo.Transaction(ctx, func(tx pgx.Tx) error {
		id, err := uc.repo.GetSheltersRepo().CreateWithConn(ctx, tx, shelter)
		if err != nil {
			return errors.Wrap(err, "failed to create shelter entity")
		}

		_, err = uc.repo.GetUsersRepo().CreateWithConn(ctx, tx, entity.User{
			Email:     email,
			Role:      entity.ManagerUserRole,
			ShelterID: id,
		})
		if err != nil {
			return errors.Wrap(err, "failed to create user entity")
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to process transaction")
	}

	return nil
}
