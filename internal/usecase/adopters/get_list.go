package adopters

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetList(ctx context.Context, filterPhoneNumber string) ([]responses.Adopter, error) {
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
