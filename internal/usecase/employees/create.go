package employees

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) Create(ctx context.Context, userId int64, req requests.CreateEmployee) (int64, error) {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get user entity")
	}
	if user == nil {
		return 0, exceptions.NewPermissionDeniedException()
	}
	if user.Role != structs.ManagerUserRole {
		return 0, exceptions.NewPermissionDeniedException()
	}

	shelter, err := uc.repo.GetSheltersRepo().Get(ctx, user.ShelterID.Int64)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get shelter entity")
	}

	var id int64
	err = uc.repo.Transaction(ctx, func(tx pgx.Tx) error {
		id, err = uc.repo.GetUsersRepo().CreateWithConn(ctx, tx, entity.User{
			Email:     req.Email,
			Role:      structs.EmployeeUserRole,
			ShelterID: user.ShelterID,
		})
		if err != nil {
			return errors.Wrap(err, "failed to create user entity")
		}

		err = uc.emailsProvider.SendInvitationEmail(shelter.Name, req.Email)
		if err != nil {
			return errors.Wrap(err, "failed to send invitation email")
		}

		return nil
	})
	if err != nil {
		return 0, errors.Wrap(err, "failed to process transaction")
	}

	return id, nil
}
