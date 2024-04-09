package shelters

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetNames(ctx context.Context, filterName string) ([]responses.ShelterName, error) {
	sheltersNames, err := uc.repo.GetSheltersRepo().SelectNames(ctx, filterName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get shelters' names entity")
	}

	return formSheltersNamesResponse(sheltersNames), nil
}

func formSheltersNamesResponse(sheltersNames []entity.SheltersNames) []responses.ShelterName {
	response := make([]responses.ShelterName, 0)
	for _, shelterName := range sheltersNames {
		response = append(response, responses.ShelterName{
			ID:   shelterName.ID,
			Name: shelterName.Name,
		})
	}

	return response
}
