package repo

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/helpers"
	"github.com/pets-shelters/backend-svc/pkg/date"
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

func (r *TasksRepo) Get(ctx context.Context, id int64) (*entity.Task, error) {
	sql, args, err := r.Builder.
		Select("*").
		From(tasksTableName).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build get task query")
	}

	var task entity.Task
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&task.ID, &task.Description, &task.StartDate, &task.EndDate, &task.Time)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to Query get task query")
	}

	return &task, nil
}

func (r *TasksRepo) SelectShelterID(ctx context.Context, taskId int64) (int64, error) {
	sql, args, err := r.Builder.
		Select(fmt.Sprintf("%s.id", sheltersTableName)).
		From(tasksTableName).
		LeftJoin(fmt.Sprintf("%s ON %s.task_id = %s.id", tasksAnimalsTableName, tasksAnimalsTableName, tasksTableName)).
		LeftJoin(fmt.Sprintf("%s ON %s.animal_id = %s.id", animalsTableName, tasksAnimalsTableName, animalsTableName)).
		LeftJoin(fmt.Sprintf("%s ON %s.location_id = %s.id", locationsTableName, animalsTableName, locationsTableName)).
		LeftJoin(fmt.Sprintf("%s ON %s.shelter_id = %s.id", sheltersTableName, locationsTableName, sheltersTableName)).
		Where(squirrel.Eq{fmt.Sprintf("%s.id", tasksTableName): taskId}).
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build select task's shelter_id query")
	}

	var shelterId int64
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&shelterId)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, errors.Wrap(err, "failed to QueryRow select task's shelter_id query")
	}

	return shelterId, nil
}

func (r *TasksRepo) Delete(ctx context.Context, id int64) (int64, error) {
	sql, args, err := r.Builder.
		Delete(tasksTableName).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build delete task query")
	}

	commandTag, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "failed to Exec delete task query")
	}

	return commandTag.RowsAffected(), nil
}

func (r *TasksRepo) SelectWithExecutions(ctx context.Context, filters entity.TasksFilters) ([]entity.TaskWithExecutions, error) {
	builder := r.Builder.
		Select(fmt.Sprintf("%s.*", tasksTableName),
			fmt.Sprintf("%s.user_id, %s.date, %s.done_at", tasksExecutionsTableName, tasksExecutionsTableName, tasksExecutionsTableName),
			fmt.Sprintf("%s.animal_id", tasksAnimalsTableName)).
		From(tasksTableName)
	builder = r.applyFilters(builder, filters)
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build select tasks with executions query")
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query select tasks with executions query")
	}

	keys := make([]int64, 0)
	tasksWithExecutions := make(map[int64]entity.TaskWithExecutions)
	defer rows.Close()
	for rows.Next() {
		var task entity.Task
		var animalId int64
		var execution entity.TaskExecutionForListNull
		err = rows.Scan(&task.ID, &task.Description, &task.StartDate, &task.EndDate, &task.Time, &execution.UserID, &execution.Date, &execution.DoneAt, &animalId)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan task with executions entity")
		}

		taskWithExecution, ok := tasksWithExecutions[task.ID]
		if !ok {
			taskWithExecution = entity.TaskWithExecutions{
				Task:       task,
				AnimalID:   animalId,
				Executions: []entity.TaskExecutionForList{},
			}
			keys = append(keys, task.ID)
		}

		if execution.DoneAt.Valid {
			taskWithExecution.Executions = append(taskWithExecution.Executions, entity.TaskExecutionForList{
				UserID: execution.UserID,
				Date:   execution.Date,
				DoneAt: execution.DoneAt.Time,
			})
		}
		tasksWithExecutions[task.ID] = taskWithExecution
	}

	return helpers.MapToArray[entity.TaskWithExecutions, int64](keys, tasksWithExecutions), nil
}

func (r *TasksRepo) applyFilters(builder squirrel.SelectBuilder, filters entity.TasksFilters) squirrel.SelectBuilder {
	if filters.Date != nil {
		builder = builder.LeftJoin(fmt.Sprintf("%s ON %s.task_id = %s.id AND %s.date = ?",
			tasksExecutionsTableName, tasksExecutionsTableName, tasksTableName, tasksExecutionsTableName), filters.Date)
	} else {
		builder = builder.LeftJoin(fmt.Sprintf("%s ON %s.task_id = %s.id",
			tasksExecutionsTableName, tasksExecutionsTableName, tasksTableName))
	}

	if filters.ShelterID != nil {
		builder = builder.LeftJoin(fmt.Sprintf("%s ON %s.task_id = %s.id", tasksAnimalsTableName, tasksAnimalsTableName, tasksTableName)).
			LeftJoin(fmt.Sprintf("%s ON %s.animal_id = %s.id", animalsTableName, tasksAnimalsTableName, animalsTableName)).
			LeftJoin(fmt.Sprintf("%s ON %s.location_id = %s.id", locationsTableName, animalsTableName, locationsTableName)).
			LeftJoin(fmt.Sprintf("%s ON %s.shelter_id = %s.id", sheltersTableName, locationsTableName, sheltersTableName)).
			Where(squirrel.Eq{fmt.Sprintf("%s.id", sheltersTableName): *filters.ShelterID})
	}
	if filters.Date != nil {
		builder = builder.Where(squirrel.LtOrEq{"start_date": *filters.Date})
		builder = builder.Where(squirrel.GtOrEq{"end_date": *filters.Date})
		builder = builder.OrderBy("time")
	}
	if filters.AnimalID != nil {
		builder = builder.Where(squirrel.Eq{fmt.Sprintf("%s.id", animalsTableName): filters.AnimalID})
	}

	return builder
}

func (r *TasksRepo) SelectForAnimal(ctx context.Context, animalId int64, pagination *entity.Pagination) ([]entity.TaskForAnimal, error) {
	builder := r.Builder.
		Select(fmt.Sprintf("%s.*", tasksTableName), fmt.Sprintf("COUNT(%s.id)", tasksExecutionsTableName)).
		From(tasksTableName).
		LeftJoin(fmt.Sprintf("%s ON %s.task_id = %s.id", tasksExecutionsTableName, tasksExecutionsTableName, tasksTableName)).
		LeftJoin(fmt.Sprintf("%s ON %s.task_id = %s.id", tasksAnimalsTableName, tasksAnimalsTableName, tasksTableName)).
		LeftJoin(fmt.Sprintf("%s ON %s.animal_id = %s.id", animalsTableName, tasksAnimalsTableName, animalsTableName)).
		Where(squirrel.Eq{fmt.Sprintf("%s.id", animalsTableName): animalId}).
		GroupBy(fmt.Sprintf("%s.id", tasksTableName))
	if pagination != nil {
		builder = helpers.ApplyPagination(builder, fmt.Sprintf("%s.start_date DESC, %s.time DESC", tasksTableName, tasksTableName), *pagination)
	}
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build select tasks for animal query")
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query select tasks for animal query")
	}

	tasksForAnimal := make([]entity.TaskForAnimal, 0)
	defer rows.Close()
	for rows.Next() {
		var taskForAnimal entity.TaskForAnimal
		err = rows.Scan(&taskForAnimal.ID, &taskForAnimal.Description, &taskForAnimal.StartDate,
			&taskForAnimal.EndDate, &taskForAnimal.Time, &taskForAnimal.ExecutionsNumber)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan task for animal entity")
		}

		tasksForAnimal = append(tasksForAnimal, taskForAnimal)
	}

	return tasksForAnimal, nil
}

// Count Reduce join level by one (use animal_id)
func (r *TasksRepo) Count(ctx context.Context, animalId int64) (int64, error) {
	sql, args, err := r.Builder.
		Select(fmt.Sprintf("COUNT(%s.*)", tasksTableName)).
		From(tasksTableName).
		LeftJoin(fmt.Sprintf("%s ON %s.task_id = %s.id", tasksExecutionsTableName, tasksExecutionsTableName, tasksTableName)).
		LeftJoin(fmt.Sprintf("%s ON %s.task_id = %s.id", tasksAnimalsTableName, tasksAnimalsTableName, tasksTableName)).
		LeftJoin(fmt.Sprintf("%s ON %s.animal_id = %s.id", animalsTableName, tasksAnimalsTableName, animalsTableName)).
		Where(squirrel.Eq{fmt.Sprintf("%s.id", animalsTableName): animalId}).
		GroupBy(fmt.Sprintf("%s.id", tasksTableName)).
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build count tasks for animal query")
	}

	var totalEntities int64
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&totalEntities)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow count tasks for animal query")
	}

	return totalEntities, nil
}

func (r *TasksRepo) SelectForEmails(ctx context.Context, date date.Date) ([]entity.EmployeeTasks, error) {
	sql, args, err := r.Builder.
		Select(fmt.Sprintf("%s.email", usersTableName),
			fmt.Sprintf("%s.description, %s.time", tasksTableName, tasksTableName),
			fmt.Sprintf("%s.name, %s.type", animalsTableName, animalsTableName)).
		From(tasksTableName).
		LeftJoin(fmt.Sprintf("%s ON %s.task_id = %s.id", tasksAnimalsTableName, tasksAnimalsTableName, tasksTableName)).
		LeftJoin(fmt.Sprintf("%s ON %s.animal_id = %s.id", animalsTableName, tasksAnimalsTableName, animalsTableName)).
		LeftJoin(fmt.Sprintf("%s ON %s.location_id = %s.id", locationsTableName, animalsTableName, locationsTableName)).
		LeftJoin(fmt.Sprintf("%s ON %s.shelter_id = %s.id", sheltersTableName, locationsTableName, sheltersTableName)).
		LeftJoin(fmt.Sprintf("%s ON %s.shelter_id = %s.id", usersTableName, usersTableName, sheltersTableName)).
		Where(squirrel.LtOrEq{"start_date": date}).
		Where(squirrel.GtOrEq{"end_date": date}).
		OrderBy("time").
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build select tasks for emails query")
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query select tasks for emails query")
	}

	keys := make([]string, 0)
	employeesTasks := make(map[string]entity.EmployeeTasks)
	defer rows.Close()
	for rows.Next() {
		var employee string
		var taskForEmail entity.TaskForEmail
		err = rows.Scan(&employee, &taskForEmail.Description, &taskForEmail.Time, &taskForEmail.AnimalName, &taskForEmail.AnimalType)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan task with executions entity")
		}

		employeeTask, ok := employeesTasks[employee]
		if !ok {
			employeeTask = entity.EmployeeTasks{
				EmployeeEmail: employee,
				Tasks:         []entity.TaskForEmail{},
			}
			keys = append(keys, employee)
		}

		employeeTask.Tasks = append(employeeTask.Tasks, taskForEmail)

		employeesTasks[employee] = employeeTask
	}

	return helpers.MapToArray[entity.EmployeeTasks, string](keys, employeesTasks), nil
}
