package adopters

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetList(ctx context.Context, userId int64, filterPhoneNumber string) ([]responses.Adopter, error) {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user entity")
	}
	if user == nil {
		return nil, exceptions.NewPermissionDeniedException()
	}

	adopters, err := uc.repo.GetAdoptersRepo().Select(ctx, filterPhoneNumber)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select adopters entities")
	}

	return formAdoptersResponse(adopters), nil
}

func formAdoptersResponse(adopters []entity.Adopter) []responses.Adopter {
	response := make([]responses.Adopter, 0)
	for _, adopter := range adopters {
		response = append(response, responses.Adopter{
			ID:          adopter.ID,
			Name:        adopter.Name,
			PhoneNumber: adopter.PhoneNumber,
		})
	}

	return response
}
