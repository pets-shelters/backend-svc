package repo

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
)

const (
	tasksExecutionsTableName = "tasks_executions"
)

type TasksExecutionsRepo struct {
	*postgres.Postgres
}

func NewTasksExecutionsRepo(pg *postgres.Postgres) *TasksExecutionsRepo {
	return &TasksExecutionsRepo{pg}
}

func (r *TasksExecutionsRepo) CreateWithConn(ctx context.Context, conn usecase.IConnection, taskExecution entity.TaskExecution) (int64, error) {
	sql, args, err := r.Builder.
		Insert(tasksExecutionsTableName).
		Columns("task_id", "user_id", "date", "done_at").
		Values(taskExecution.TaskID, taskExecution.UserID, taskExecution.Date, taskExecution.DoneAt).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build task_execution insert query")
	}

	var id int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow task_execution insert query")
	}

	return id, nil
}

func (r *TasksExecutionsRepo) Create(ctx context.Context, taskExecution entity.TaskExecution) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, taskExecution)
}
