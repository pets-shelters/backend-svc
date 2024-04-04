package shelters

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) Update(ctx context.Context, req requests.UpdateShelter, userId int64, shelterId int64) error {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to get user entity")
	}
	if user.ShelterID.Int64 != shelterId {
		return exceptions.NewPermissionDeniedException()
	}
	if user.Role != structs.ManagerUserRole {
		return exceptions.NewPermissionDeniedException()
	}

	err = uc.repo.Transaction(ctx, func(tx pgx.Tx) error {
		if req.Logo != nil {
			tempFile, err := uc.repo.GetTemporaryFilesRepo().DeleteWithConn(ctx, tx, *req.Logo)
			if err != nil {
				return errors.Wrap(err, "failed to delete temporary_file entity")
			}
			if tempFile == nil {
				return exceptions.NewFileNotFoundException()
			}
			if tempFile.UserID != userId {
				return exceptions.NewPermissionDeniedException()
			}
		}

		rowsAffected, err := uc.repo.GetSheltersRepo().Update(ctx, tx, shelterId, entity.UpdateShelter{
			Name:        req.Name,
			Logo:        req.Logo,
			PhoneNumber: req.PhoneNumber,
			Instagram:   req.Instagram,
			Facebook:    req.Facebook,
		})
		if err != nil {
			return errors.Wrap(err, "failed to update shelter entity")
		}
		if rowsAffected == 0 {
			return exceptions.NewNotFoundException()
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to process transaction")
	}

	return nil
}
