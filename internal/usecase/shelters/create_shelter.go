package shelters

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/usecase/postgres/entity"
	"github.com/pkg/errors"
	"time"
)

func (uc *UseCase) Create(ctx context.Context, req requests.CreateShelter, userEmail string) error {
	err := uc.repo.Transaction(ctx, func(tx pgx.Tx) error {
		id, err := uc.repo.GetSheltersRepo().CreateWithConn(ctx, tx, entity.Shelter{
			Name:        req.Name,
			Logo:        req.Logo,
			City:        req.City,
			PhoneNumber: req.PhoneNumber,
			Instagram:   req.Instagram,
			Facebook:    req.Facebook,
			CreatedAt:   time.Now(),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create shelter entity")
		}

		rowsAffected, err := uc.repo.GetUsersRepo().UpdateShelterIDWithConn(ctx, tx, userEmail, id)
		if err != nil {
			return errors.Wrap(err, "failed to update user entity")
		}
		if rowsAffected == 0 {
			return exceptions.NewUserHasShelterException()
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to process transaction")
	}

	return nil
}
