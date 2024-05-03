package tasks

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) Create(ctx context.Context, req requests.CreateTask, userId int64) (int64, error) {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get user entity")
	}
	if user == nil {
		return 0, exceptions.NewPermissionDeniedException()
	}

	shelterId, err := uc.repo.GetAnimalsRepo().SelectShelterID(ctx, req.AnimalID)
	if err != nil {
		return 0, errors.Wrap(err, "failed to select animal's shelter_id")
	}
	if shelterId == 0 {
		return 0, exceptions.NewNotFoundException()
	}
	if user.ShelterID.Int64 != shelterId {
		return 0, exceptions.NewPermissionDeniedException()
	}

	var id int64
	err = uc.repo.Transaction(ctx, func(tx pgx.Tx) error {
		id, err = uc.repo.GetTasksRepo().CreateWithConn(ctx, tx, entity.Task{
			Description: req.Description,
			StartDate:   req.StartDate,
			EndDate:     req.EndDate,
			Time:        req.Time,
		})
		if err != nil {
			return errors.Wrap(err, "failed to create task entity")
		}

		_, err = uc.repo.GetTasksAnimalsRepo().CreateWithConn(ctx, tx, entity.TaskAnimal{
			AnimalID: req.AnimalID,
			TaskID:   id,
		})
		if err != nil {
			return errors.Wrap(err, "failed to create task_animal entity")
		}

		return nil
	})
	if err != nil {
		return 0, errors.Wrap(err, "failed to process transaction")
	}

	return id, nil
}
