package repo

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	entity "github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
)

const sheltersTableName = "shelters"

type SheltersRepo struct {
	*postgres.Postgres
}

func NewSheltersRepo(pg *postgres.Postgres) *SheltersRepo {
	return &SheltersRepo{pg}
}

func (r *SheltersRepo) CreateWithConn(ctx context.Context, conn usecase.IConnection, shelter entity.Shelter) (int64, error) {
	sql, args, err := r.Builder.
		Insert(sheltersTableName).
		Columns("name", "logo", "phone_number", "instagram", "facebook", "created_at").
		Values(shelter.Name, shelter.Logo, shelter.PhoneNumber, shelter.Instagram, shelter.Facebook, shelter.CreatedAt).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build shelter insert query")
	}

	var id int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow shelter insert query")
	}

	return id, nil
}

func (r *SheltersRepo) Create(ctx context.Context, shelter entity.Shelter) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, shelter)
}

func (r *SheltersRepo) SelectWithConn(ctx context.Context, conn usecase.IConnection) ([]entity.Shelter, error) {
	sql, _, err := r.Builder.
		Select("*").
		From(sheltersTableName).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build select shelters query")
	}

	rows, err := conn.Query(ctx, sql)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query select shelters query")
	}

	var shelters []entity.Shelter
	defer rows.Close()
	for rows.Next() {
		var shelter entity.Shelter
		err = rows.Scan(&shelter.ID, &shelter.Logo, &shelter.Name, &shelter.PhoneNumber,
			&shelter.CreatedAt, &shelter.Instagram, &shelter.Facebook)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan shelter entity")
		}

		shelters = append(shelters, shelter)
	}

	return shelters, nil
}

func (r *SheltersRepo) Select(ctx context.Context) ([]entity.Shelter, error) {
	return r.SelectWithConn(ctx, r.Pool)
}

func (r *SheltersRepo) Get(ctx context.Context, id int64) (*entity.Shelter, error) {
	sql, args, err := r.Builder.
		Select("*").
		From(sheltersTableName).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build get shelter query")
	}

	var shelter entity.Shelter
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&shelter.ID, &shelter.Name, &shelter.Logo, &shelter.PhoneNumber,
		&shelter.Instagram, &shelter.Facebook, &shelter.CreatedAt)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to Query get shelter query")
	}

	return &shelter, nil
}

func (r *SheltersRepo) Update(ctx context.Context, conn usecase.IConnection, id int64, updateParams entity.UpdateShelter) (int64, error) {
	sql, args, err := r.applyUpdateParams(updateParams).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build update shelter query")
	}

	commandTag, err := conn.Exec(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "failed to Query select employees query")
	}

	return commandTag.RowsAffected(), nil
}

func (r *SheltersRepo) applyUpdateParams(updateParams entity.UpdateShelter) squirrel.UpdateBuilder {
	builder := r.Builder.Update(sheltersTableName)
	if updateParams.Name != nil {
		builder = builder.Set("name", *updateParams.Name)
	}
	if updateParams.Logo != nil {
		builder = builder.Set("logo", *updateParams.Logo)
	}
	if updateParams.PhoneNumber != nil {
		builder = builder.Set("phone_number", *updateParams.PhoneNumber)
	}
	if updateParams.Instagram != nil {
		builder = builder.Set("instagram", *updateParams.Instagram)
	}
	if updateParams.Facebook != nil {
		builder = builder.Set("facebook", *updateParams.Facebook)
	}

	return builder
}

func (r *SheltersRepo) GetNames(ctx context.Context, filterName string) ([]string, error) {
	sql, args, err := r.Builder.
		Select("name").
		From(sheltersTableName).
		Where(squirrel.Like{"name": "%" + filterName + "%"}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build get shelters' names query")
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query select shelters' names query")
	}

	sheltersNames := make([]string, 0)
	defer rows.Close()
	for rows.Next() {
		var shelterName string
		err = rows.Scan(&shelterName)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan shelter' name entity")
		}

		sheltersNames = append(sheltersNames, shelterName)
	}

	return sheltersNames, nil
}
