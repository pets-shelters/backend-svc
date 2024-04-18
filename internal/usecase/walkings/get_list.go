package walkings

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetList(ctx context.Context, reqFilters requests.WalkingsFilters, reqPagination *requests.Pagination, userId int64) ([]responses.Walking, *responses.PaginationMetadata, error) {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get user entity")
	}
	if user == nil {
		return nil, nil, exceptions.NewPermissionDeniedException()
	}

	var pagination *entity.Pagination
	if reqPagination != nil {
		pagination = &entity.Pagination{
			Page:  uint64(*reqPagination.Page),
			Limit: uint64(*reqPagination.Limit),
		}
	}

	filters := entity.WalkingsFilters{
		ShelterID: &user.ShelterID.Int64,
		AnimalID:  reqFilters.AnimalID,
		Status:    reqFilters.Status,
		Date:      reqFilters.Date,
	}

	walkings, err := uc.repo.GetWalkingsRepo().Select(ctx, filters, pagination)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to select walkings entities")
	}

	var paginationMetadata *responses.PaginationMetadata
	if reqPagination != nil {
		totalEntities, err := uc.repo.GetWalkingsRepo().Count(ctx, filters)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to count animals entities")
		}
		paginationMetadata = &responses.PaginationMetadata{
			CurrentPage:  *reqPagination.Page,
			TotalRecords: totalEntities,
		}
	}
	return formWalkingsResponse(walkings), paginationMetadata, nil
}

func formWalkingsResponse(walkings []entity.Walking) []responses.Walking {
	response := make([]responses.Walking, 0)
	for _, walking := range walkings {
		response = append(response, responses.Walking{
			ID:          walking.ID,
			Status:      walking.Status,
			AnimalID:    walking.AnimalID,
			Name:        walking.Name,
			PhoneNumber: walking.PhoneNumber,
			Date:        walking.Date,
			Time:        walking.Time,
		})
	}

	return response
}
