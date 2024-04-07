package locations

import (
	"context"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetCities(ctx context.Context) ([]string, error) {
	cities, err := uc.repo.GetLocationsRepo().SelectUniqueCities(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select locations' unique cities entities")
	}

	return cities, nil
}
