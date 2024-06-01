package adopters

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
	"log"
)

func (uc *UseCase) Create(ctx context.Context, req requests.CreateAdopter) (int64, error) {
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
		return 0, errors.Wrap(err, "failed to process transaction")
	}

	log.Printf("usecase %+v", id)
	return id, nil
}
