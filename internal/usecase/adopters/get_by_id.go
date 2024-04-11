package adopters

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetById(ctx context.Context, userId int64, adopterId int64) (*responses.Adopter, error) {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user entity")
	}
	if user == nil {
		return nil, exceptions.NewPermissionDeniedException()
	}

	adopter, err := uc.repo.GetAdoptersRepo().Get(ctx, adopterId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get adopter entity")
	}
	if adopter == nil {
		return nil, exceptions.NewNotFoundException()
	}

	return &responses.Adopter{
		ID:          adopter.ID,
		Name:        adopter.Name,
		PhoneNumber: adopter.PhoneNumber,
	}, nil
}
