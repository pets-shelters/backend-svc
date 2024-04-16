package tasks

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetListForAnimal(ctx context.Context, userId int64, animalId int64) ([]responses.TaskForAnimal, error) {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user entity")
	}

	shelterId, err := uc.repo.GetAnimalsRepo().SelectShelterID(ctx, animalId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select animal's shelter_id")
	}
	if shelterId == 0 {
		return nil, exceptions.NewNotFoundException()
	}
	if user.ShelterID.Int64 != shelterId {
		return nil, exceptions.NewPermissionDeniedException()
	}

	tasksForAnimal, err := uc.repo.GetTasksRepo().SelectForAnimal(ctx, animalId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select tasks entities")
	}

	return formTasksForAnimalResponse(tasksForAnimal), nil
}

func formTasksForAnimalResponse(tasksForAnimal []entity.TaskForAnimal) []responses.TaskForAnimal {
	response := make([]responses.TaskForAnimal, 0)
	for _, taskForAnimal := range tasksForAnimal {
		responseTaskForAnimal := responses.TaskForAnimal{
			ID:               taskForAnimal.ID,
			Description:      taskForAnimal.Description,
			StartDate:        taskForAnimal.StartDate,
			EndDate:          taskForAnimal.EndDate,
			Time:             taskForAnimal.Time,
			ExecutionsNumber: taskForAnimal.ExecutionsNumber,
		}
		response = append(response, responseTaskForAnimal)
	}

	return response
}
