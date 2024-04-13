package animals

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) Update(ctx context.Context, req requests.UpdateAnimal, userId int64, animalId int64) error {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to get user entity")
	}
	if user == nil {
		return exceptions.NewPermissionDeniedException()
	}

	var newLocation *entity.Location
	if req.LocationID != nil {
		newLocation, err = uc.repo.GetLocationsRepo().Get(ctx, *req.LocationID)
		if err != nil {
			return errors.Wrap(err, "failed to get new location entity")
		}
		if newLocation == nil {
			return exceptions.NewLocationNotFoundException()
		}
	}

	shelterId, err := uc.repo.GetAnimalsRepo().SelectShelterID(ctx, animalId)
	if err != nil {
		return errors.Wrap(err, "failed to select animal's shelter_id")
	}
	if user.ShelterID.Int64 != shelterId {
		return exceptions.NewPermissionDeniedException()
	}
	if newLocation != nil && shelterId != newLocation.ShelterID {
		return exceptions.NewPermissionDeniedException()
	}

	err = uc.repo.Transaction(ctx, func(tx pgx.Tx) error {
		if req.Photo != nil {
			tempFile, err := uc.repo.GetTemporaryFilesRepo().DeleteWithConn(ctx, tx, *req.Photo)
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

		rowsAffected, err := uc.repo.GetAnimalsRepo().Update(ctx, tx, animalId, entity.UpdateAnimal{
			LocationID:         req.LocationID,
			Photo:              req.Photo,
			Sterilized:         req.Sterilized,
			AdopterID:          req.AdopterID,
			PublicDescription:  req.PublicDescription,
			PrivateDescription: req.PrivateDescription,
		})
		if err != nil {
			return errors.Wrap(err, "failed to update animal entity")
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
