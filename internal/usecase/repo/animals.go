package repo

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/helpers"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
	"log"
)

const (
	animalsTableName = "animals"
)

type AnimalsRepo struct {
	*postgres.Postgres
}

func NewAnimalsRepo(pg *postgres.Postgres) *AnimalsRepo {
	return &AnimalsRepo{pg}
}

func (r *AnimalsRepo) CreateWithConn(ctx context.Context, conn usecase.IConnection, animal entity.Animal) (int64, error) {
	sql, args, err := r.Builder.
		Insert(animalsTableName).
		Columns("location_id", "photo", "name", "birth_date", "type", "gender", "sterilized",
			"private_description", "public_description").
		Values(animal.LocationID, animal.Photo, animal.Name, animal.BirthDate, animal.Type, animal.Gender, animal.Sterilized,
			animal.PrivateDescription, animal.PublicDescription).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build animal insert query")
	}

	var id int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow animal insert query")
	}

	return id, nil
}

func (r *AnimalsRepo) Create(ctx context.Context, animal entity.Animal) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, animal)
}

func (r *AnimalsRepo) Select(ctx context.Context, filters entity.AnimalsFilters, pagination *entity.Pagination) ([]entity.AnimalForList, error) {
	builder := r.Builder.
		Select(fmt.Sprintf("%s.id", animalsTableName), "name", "birth_date", "type",
			"bucket", "path").
		From(animalsTableName)
	builder = r.applyFilters(builder, filters)
	if pagination != nil {
		builder = helpers.ApplyPagination(builder, animalsTableName, *pagination)
	}
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build select animals query")
	}

	log.Printf("%+v", sql)
	log.Printf("%+v", args)
	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query select animals query")
	}

	animals := make([]entity.AnimalForList, 0)
	defer rows.Close()
	for rows.Next() {
		var animal entity.AnimalForList
		err = rows.Scan(&animal.ID, &animal.Name, &animal.BirthDate, &animal.Type, &animal.PhotoBucket, &animal.PhotoPath)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan animal entity")
		}

		animals = append(animals, animal)
	}

	return animals, nil
}

func (r *AnimalsRepo) applyFilters(builder squirrel.SelectBuilder, filters entity.AnimalsFilters) squirrel.SelectBuilder {
	builder = builder.LeftJoin(fmt.Sprintf("%s ON %s.photo = %s.id", filesTableName, animalsTableName, filesTableName))
	if filters.ShelterID != nil || filters.City != nil {
		builder = builder.LeftJoin(fmt.Sprintf("%s ON %s.location_id = %s.id", locationsTableName, animalsTableName, locationsTableName))
	}

	if filters.ShelterID != nil {
		builder = builder.Where(squirrel.Eq{fmt.Sprintf("%s.shelter_id", locationsTableName): filters.ShelterID})
	}
	if filters.LocationID != nil {
		builder = builder.Where(squirrel.Eq{"location_id": filters.LocationID})
	}
	if filters.Gender != nil {
		builder = builder.Where(squirrel.Eq{"gender": *filters.Gender})
	}
	if filters.Sterilized != nil {
		builder = builder.Where(squirrel.Eq{"sterilized": *filters.Sterilized})
	}
	if filters.BirthDateFrom != nil {
		builder = builder.Where(squirrel.GtOrEq{"birth_date": *filters.BirthDateFrom})
	}
	if filters.BirthDateTo != nil {
		builder = builder.Where(squirrel.LtOrEq{"birth_date": *filters.BirthDateTo})
	}
	if filters.Type != nil {
		builder = builder.Where(squirrel.Eq{"type": filters.Type})
	}
	if filters.Name != nil {
		builder = builder.Where(squirrel.Like{"name": "%" + *filters.Name + "%"})
	}
	if filters.City != nil {
		builder = builder.Where(squirrel.Eq{fmt.Sprintf("%s.city", locationsTableName): filters.City})
	}

	return builder
}

func (r *AnimalsRepo) Count(ctx context.Context, filters entity.AnimalsFilters) (int64, error) {
	builder := r.Builder.
		Select(fmt.Sprintf("COUNT(%s.*)", animalsTableName)).
		From(animalsTableName)
	builder = r.applyFilters(builder, filters)
	sql, args, err := builder.ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build count animals query")
	}

	log.Printf("%+v", sql)
	log.Printf("%+v", args)
	var totalEntities int64
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&totalEntities)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow count animals query")
	}

	return totalEntities, nil
}
