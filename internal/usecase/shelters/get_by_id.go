package shelters

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetById(ctx context.Context, shelterId int64) (*responses.Shelter, error) {
	shelter, err := uc.repo.GetSheltersRepo().Get(ctx, shelterId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get shelter entity")
	}
	if shelter == nil {
		return nil, exceptions.NewNotFoundException()
	}

	file, err := uc.repo.GetFilesRepo().Get(ctx, shelter.Logo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get file entity")
	}

	return &responses.Shelter{
		ID:          shelter.ID,
		Name:        shelter.Name,
		Logo:        uc.s3Endpoint + "/" + file.Bucket + file.Path,
		PhoneNumber: shelter.PhoneNumber,
		CreatedAt:   shelter.CreatedAt,
		Instagram:   shelter.Instagram,
		Facebook:    shelter.Facebook,
	}, nil
}
