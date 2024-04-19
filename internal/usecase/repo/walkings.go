package repo

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/helpers"
	"github.com/pets-shelters/backend-svc/pkg/date"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
)

const (
	walkingsTableName = "walkings"
)

type WalkingsRepo struct {
	*postgres.Postgres
}

func NewWalkingsRepo(pg *postgres.Postgres) *WalkingsRepo {
	return &WalkingsRepo{pg}
}

func (r *WalkingsRepo) CreateWithConn(ctx context.Context, conn usecase.IConnection, walking entity.Walking) (int64, error) {
	sql, args, err := r.Builder.
		Insert(walkingsTableName).
		Columns("status", "animal_id", "name", "phone_number", "date", "time").
		Values(walking.Status, walking.AnimalID, walking.Name, walking.PhoneNumber, walking.Date, walking.Time).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build walking insert query")
	}

	var id int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow walking insert query")
	}

	return id, nil
}

func (r *WalkingsRepo) Create(ctx context.Context, walking entity.Walking) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, walking)
}

func (r *WalkingsRepo) Select(ctx context.Context, filters entity.WalkingsFilters, pagination *entity.Pagination) ([]entity.Walking, error) {
	builder := r.Builder.
		Select(fmt.Sprintf("%s.id, %s.status, %s.animal_id, %s.name, %s.phone_number, %s.date, %s.time",
			walkingsTableName, walkingsTableName, walkingsTableName, walkingsTableName, walkingsTableName, walkingsTableName, walkingsTableName)).
		From(walkingsTableName)
	builder = r.applyFilters(builder, filters)
	if pagination != nil {
		builder = helpers.ApplyPagination(builder, "date,time", walkingsTableName, *pagination)
	}
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build select walkings query")
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query select walkings query")
	}

	walkings := make([]entity.Walking, 0)
	defer rows.Close()
	for rows.Next() {
		var walking entity.Walking
		err = rows.Scan(&walking.ID, &walking.Status, &walking.AnimalID, &walking.Name,
			&walking.PhoneNumber, &walking.Date, &walking.Time)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan walking entity")
		}

		walkings = append(walkings, walking)
	}

	return walkings, nil
}

func (r *WalkingsRepo) applyFilters(builder squirrel.SelectBuilder, filters entity.WalkingsFilters) squirrel.SelectBuilder {
	if filters.ShelterID != nil {
		builder = builder.LeftJoin(fmt.Sprintf("%s ON %s.animal_id = %s.id", animalsTableName, walkingsTableName, animalsTableName)).
			LeftJoin(fmt.Sprintf("%s ON %s.location_id = %s.id", locationsTableName, animalsTableName, locationsTableName)).
			LeftJoin(fmt.Sprintf("%s ON %s.shelter_id = %s.id", sheltersTableName, locationsTableName, sheltersTableName)).
			Where(squirrel.Eq{fmt.Sprintf("%s.id", sheltersTableName): *filters.ShelterID})
	}
	if filters.AnimalID != nil {
		builder = builder.Where(squirrel.Eq{fmt.Sprintf("%s.animal_id", walkingsTableName): *filters.AnimalID})
	}
	if filters.Date != nil {
		builder = builder.Where(squirrel.Eq{fmt.Sprintf("%s.date", walkingsTableName): *filters.Date})
	}
	if filters.Status != nil {
		builder = builder.Where(squirrel.Eq{fmt.Sprintf("%s.status", walkingsTableName): *filters.Status})
	}

	return builder
}

func (r *WalkingsRepo) DeleteWithConn(ctx context.Context, conn usecase.IConnection, id int64) (*entity.Walking, error) {
	sql, args, err := r.Builder.
		Delete(walkingsTableName).
		Where(squirrel.Eq{"id": id}).
		Suffix("returning *").
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build delete walking query")
	}

	var walking entity.Walking
	err = conn.QueryRow(ctx, sql, args...).Scan(&walking.ID, &walking.Status, &walking.AnimalID, &walking.Name,
		&walking.PhoneNumber, &walking.Date, &walking.Time)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to QueryRow delete walking query")
	}

	return &walking, nil
}

func (r *WalkingsRepo) Update(ctx context.Context, conn usecase.IConnection, id int64, updateParams entity.UpdateWalking) (animalId int64, err error) {
	sql, args, err := r.applyUpdateParams(updateParams).
		Where(squirrel.Eq{"id": id}).
		Suffix("returning animal_id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build update walking query")
	}

	err = conn.QueryRow(ctx, sql, args...).Scan(&animalId)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow update walking query")
	}

	return
}

func (r *WalkingsRepo) applyUpdateParams(updateParams entity.UpdateWalking) squirrel.UpdateBuilder {
	builder := r.Builder.Update(walkingsTableName)

	if updateParams.Status != nil {
		builder = builder.Set("status", *updateParams.Status)
	}
	if updateParams.Date != nil {
		builder = builder.Set("date", *updateParams.Date)
	}
	if updateParams.Time != nil {
		builder = builder.Set("time", *updateParams.Time)
	}

	return builder
}

func (r *WalkingsRepo) Count(ctx context.Context, filters entity.WalkingsFilters) (int64, error) {
	builder := r.Builder.
		Select(fmt.Sprintf("COUNT(%s.*)", walkingsTableName)).
		From(walkingsTableName)
	builder = r.applyFilters(builder, filters)
	sql, args, err := builder.ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build count walkings query")
	}

	var totalEntities int64
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&totalEntities)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow count walkings query")
	}

	return totalEntities, nil
}

func (r *WalkingsRepo) SelectForReminders(ctx context.Context, date date.Date) ([]entity.WalkingReminder, error) {
	sql, args, err := r.Builder.
		Select(fmt.Sprintf("%s.name", sheltersTableName),
			fmt.Sprintf("%s.phone_number, %s.time", walkingsTableName, walkingsTableName),
			fmt.Sprintf("%s.name, %s.type", animalsTableName, animalsTableName)).
		From(walkingsTableName).
		LeftJoin(fmt.Sprintf("%s ON %s.animal_id = %s.id", animalsTableName, walkingsTableName, animalsTableName)).
		LeftJoin(fmt.Sprintf("%s ON %s.location_id = %s.id", locationsTableName, animalsTableName, locationsTableName)).
		LeftJoin(fmt.Sprintf("%s ON %s.shelter_id = %s.id", sheltersTableName, locationsTableName, sheltersTableName)).
		Where(squirrel.Eq{"date": date}).
		Where(squirrel.Eq{"status": structs.ApprovedWalkingStatus}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build select walkings for reminders query")
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query select walkings for reminders query")
	}

	walkings := make([]entity.WalkingReminder, 0)
	defer rows.Close()
	for rows.Next() {
		var walking entity.WalkingReminder
		err = rows.Scan(&walking.ShelterName, &walking.PhoneNumber, &walking.Time, &walking.AnimalName, &walking.AnimalType)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan walking entity")
		}

		walkings = append(walkings, walking)
	}

	return walkings, nil
}
