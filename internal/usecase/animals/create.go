package animals

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) Create(ctx context.Context, req requests.CreateAnimal, userId int64) error {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to get user entity")
	}
	if user == nil {
		return exceptions.NewPermissionDeniedException()
	}

	location, err := uc.repo.GetLocationsRepo().Get(ctx, req.LocationID)
	if err != nil {
		return errors.Wrap(err, "failed to get location entity")
	}
	if location == nil {
		return exceptions.NewLocationNotFoundException()
	}

	if user.ShelterID.Int64 != location.ShelterID {
		return exceptions.NewPermissionDeniedException()
	}

	err = uc.repo.GetAnimalTypesEnumRepo().Create(ctx, req.Type)
	if err != nil {
		return errors.Wrap(err, "failed to create animal_type value")
	}

	err = uc.repo.Transaction(ctx, func(tx pgx.Tx) error {
		tempFile, err := uc.repo.GetTemporaryFilesRepo().DeleteWithConn(ctx, tx, req.Photo)
		if err != nil {
			return errors.Wrap(err, "failed to delete temporary_file entity")
		}
		if tempFile == nil {
			return exceptions.NewFileNotFoundException()
		}
		if tempFile.UserID != userId {
			return exceptions.NewPermissionDeniedException()
		}

		_, err = uc.repo.GetAnimalsRepo().CreateWithConn(ctx, tx, entity.Animal{
			Name:               req.Name,
			LocationID:         req.LocationID,
			Photo:              req.Photo,
			Gender:             req.Gender,
			Sterilized:         req.Sterilized,
			ForAdoption:        req.ForAdoption,
			ForWalking:         req.ForWalking,
			Type:               req.Type,
			BirthDate:          req.BirthDate,
			PrivateDescription: req.PrivateDescription,
			PublicDescription:  req.PublicDescription,
		})
		if err != nil {
			return errors.Wrap(err, "failed to create animal entity")
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to process transaction")
	}

	return nil
}
