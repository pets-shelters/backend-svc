package shelters

import (
	"context"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetNames(ctx context.Context, filterName string) ([]string, error) {
	sheltersNames, err := uc.repo.GetSheltersRepo().GetNames(ctx, filterName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get shelters' names entity")
	}

	return sheltersNames, nil
}
