package repo

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
)

const (
	tasksTableName = "tasks"
)

type TasksRepo struct {
	*postgres.Postgres
}

func NewTasksRepo(pg *postgres.Postgres) *TasksRepo {
	return &TasksRepo{pg}
}

func (r *TasksRepo) CreateWithConn(ctx context.Context, conn usecase.IConnection, task entity.Task) (int64, error) {
	sql, args, err := r.Builder.
		Insert(tasksTableName).
		Columns("description", "start_date", "end_date", "time").
		Values(task.Description, task.StartDate, task.EndDate, task.Time).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build task insert query")
	}

	var id int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow task insert query")
	}

	return id, nil
}

func (r *TasksRepo) Create(ctx context.Context, task entity.Task) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, task)
}
