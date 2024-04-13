package animals

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetList(ctx context.Context, reqFilters requests.AnimalsFilters, reqPagination *requests.Pagination) ([]responses.AnimalForList, *responses.PaginationMetadata, error) {
	var pagination *entity.Pagination
	if reqPagination != nil {
		pagination = &entity.Pagination{
			Page:  uint64(*reqPagination.Page),
			Limit: uint64(*reqPagination.Limit),
		}
	}

	filters := entity.AnimalsFilters{
		ShelterID:     reqFilters.ShelterID,
		LocationID:    reqFilters.LocationID,
		Gender:        reqFilters.Gender,
		Sterilized:    reqFilters.Sterilized,
		Adopted:       reqFilters.Adopted,
		BirthDateFrom: reqFilters.BirthDateFrom,
		BirthDateTo:   reqFilters.BirthDateTo,
		Type:          reqFilters.Type,
		Name:          reqFilters.Name,
		City:          reqFilters.City,
	}

	animals, err := uc.repo.GetAnimalsRepo().Select(ctx, filters, pagination)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to select animals entities")
	}

	var paginationMetadata *responses.PaginationMetadata
	if reqPagination != nil {
		totalEntities, err := uc.repo.GetAnimalsRepo().Count(ctx, filters)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to count animals entities")
		}
		paginationMetadata = &responses.PaginationMetadata{
			CurrentPage:  *reqPagination.Page,
			TotalRecords: totalEntities,
		}
	}
	return formAnimalsResponse(animals, uc.s3Endpoint), paginationMetadata, nil
}

func formAnimalsResponse(animals []entity.AnimalForList, s3Endpoint string) []responses.AnimalForList {
	response := make([]responses.AnimalForList, 0)
	for _, animal := range animals {
		response = append(response, responses.AnimalForList{
			ID:        animal.ID,
			Photo:     s3Endpoint + "/" + animal.PhotoBucket + animal.PhotoPath,
			Name:      animal.Name,
			BirthDate: animal.BirthDate,
			Type:      animal.Type,
		})
	}

	return response
}
