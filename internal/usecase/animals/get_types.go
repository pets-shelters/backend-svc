package animals

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetTypes(ctx context.Context) (*responses.AnimalTypes, error) {
	animalTypes, err := uc.repo.GetAnimalTypesEnumRepo().Select(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select animal types")
	}

	return &responses.AnimalTypes{
		AnimalTypes: animalTypes,
	}, nil
}
