package tasks

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
	"time"
)

func (uc *UseCase) SetExecution(ctx context.Context, req requests.SetTaskDone, taskId int64, userId int64) error {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to get user entity")
	}

	task, err := uc.repo.GetTasksRepo().Get(ctx, taskId)
	if err != nil {
		return errors.Wrap(err, "failed to get task entity")
	}
	if task == nil {
		return exceptions.NewNotFoundException()
	}
	if time.Time(req.Date).Before(time.Time(task.StartDate)) || time.Time(req.Date).After(time.Time(task.EndDate)) {
		return exceptions.NewInvalidTaskExecutionDateException()
	}

	taskShelterId, err := uc.repo.GetTasksRepo().SelectShelterID(ctx, taskId)
	if err != nil {
		return errors.Wrap(err, "failed to get task's shelter_id")
	}
	if taskShelterId != user.ShelterID.Int64 {
		return exceptions.NewPermissionDeniedException()
	}

	err = uc.repo.Transaction(ctx, func(tx pgx.Tx) error {
		_, err := uc.repo.GetTasksExecutionsRepo().CreateWithConn(ctx, tx, entity.TaskExecution{
			TaskID: taskId,
			UserID: sql.NullInt64{Int64: userId},
			Date:   req.Date,
			DoneAt: time.Now().UTC(),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create task entity")
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to process transaction")
	}

	return nil
}
