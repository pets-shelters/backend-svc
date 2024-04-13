package repo

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
)

const (
	tasksAnimalsTableName = "tasks_animals"
)

type TasksAnimalsRepo struct {
	*postgres.Postgres
}

func NewTasksAnimalsRepo(pg *postgres.Postgres) *TasksAnimalsRepo {
	return &TasksAnimalsRepo{pg}
}

func (r *TasksAnimalsRepo) CreateWithConn(ctx context.Context, conn usecase.IConnection, taskAnimal entity.TaskAnimal) (int64, error) {
	sql, args, err := r.Builder.
		Insert(tasksAnimalsTableName).
		Columns("animal_id", "task_id").
		Values(taskAnimal.AnimalID, taskAnimal.TaskID).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build task_animal insert query")
	}

	var id int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow task_animal insert query")
	}

	return id, nil
}

func (r *TasksAnimalsRepo) Create(ctx context.Context, taskAnimal entity.TaskAnimal) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, taskAnimal)
}
