package locations

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) Create(ctx context.Context, userId int64, req requests.CreateLocation) error {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to get user entity")
	}
	if user == nil {
		return exceptions.NewPermissionDeniedException()
	}
	if user.Role != structs.ManagerUserRole {
		return exceptions.NewPermissionDeniedException()
	}

	_, err = uc.repo.GetLocationsRepo().Create(ctx, entity.Location{
		City:      req.City,
		Address:   req.Address,
		ShelterID: user.ShelterID.Int64,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create location entity")
	}

	return nil
}
