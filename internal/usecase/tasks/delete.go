package tasks

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pkg/errors"
)

func (uc *UseCase) Delete(ctx context.Context, userId int64, taskId int64) error {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to get user entity")
	}

	taskShelterId, err := uc.repo.GetTasksRepo().SelectShelterID(ctx, taskId)
	if err != nil {
		return errors.Wrap(err, "failed to get task's shelter_id")
	}
	if taskShelterId == 0 {
		return exceptions.NewNotFoundException()
	}
	if taskShelterId != user.ShelterID.Int64 {
		return exceptions.NewPermissionDeniedException()
	}

	_, err = uc.repo.GetTasksRepo().Delete(ctx, taskId)
	if err != nil {
		return errors.Wrap(err, "failed to delete task by id")
	}

	return nil
}
