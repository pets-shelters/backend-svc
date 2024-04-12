package adopters

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) Create(ctx context.Context, req requests.CreateAdopter) (*responses.AdopterCreated, error) {
	var id int64
	var err error
	err = uc.repo.Transaction(ctx, func(tx pgx.Tx) error {
		id, err = uc.repo.GetAdoptersRepo().CreateWithConn(ctx, tx, entity.Adopter{
			Name:        req.Name,
			PhoneNumber: req.PhoneNumber,
		})
		if err != nil {
			return errors.Wrap(err, "failed to create adopter entity")
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to process transaction")
	}

	return &responses.AdopterCreated{
		AdopterID: id,
	}, nil
}
