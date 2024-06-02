package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/helpers"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
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

func (r *AnimalsRepo) CreateWithConn(ctx context.Context, conn usecase.IConnection, animal entity.CreateAnimal) (int64, error) {
	sql, args, err := r.Builder.
		Insert(animalsTableName).
		Columns("location_id", "photo", "name", "birth_date", "type", "gender", "sterilized",
			"for_adoption", "for_walking", "private_description", "public_description").
		Values(animal.LocationID, animal.Photo, animal.Name, animal.BirthDate, animal.Type, animal.Gender, animal.Sterilized,
			animal.ForAdoption, animal.ForWalking, animal.PrivateDescription, animal.PublicDescription).
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

func (r *AnimalsRepo) Create(ctx context.Context, animal entity.CreateAnimal) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, animal)
}

func (r *AnimalsRepo) Select(ctx context.Context, filters entity.AnimalsFilters, pagination *entity.Pagination) ([]entity.AnimalForList, error) {
	builder := r.Builder.
		Select(fmt.Sprintf("%s.id", animalsTableName), "name", "birth_date", "type",
			"bucket", "path").
		From(animalsTableName)
	builder = r.applyFilters(builder, filters)
	if pagination != nil {
		builder = helpers.ApplyPagination(builder, fmt.Sprintf("%s.id", animalsTableName), *pagination)
	}
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build select animals query")
	}

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
	if filters.ForAdoption != nil {
		builder = builder.Where(squirrel.Eq{"for_adoption": *filters.ForAdoption})
	}
	if filters.ForWalking != nil {
		builder = builder.Where(squirrel.Eq{"for_walking": *filters.ForWalking})
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
	if filters.Adopted != nil {
		if *filters.Adopted {
			builder = builder.Where(squirrel.NotEq{"adopter_id": sql.NullInt64{}})
		} else {
			builder = builder.Where(squirrel.Eq{"adopter_id": sql.NullInt64{}})
		}
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

	var totalEntities int64
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&totalEntities)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow count animals query")
	}

	return totalEntities, nil
}

func (r *AnimalsRepo) UpdateWithConn(ctx context.Context, conn usecase.IConnection, id int64, updateParams entity.UpdateAnimal) (int64, error) {
	sql, args, err := r.applyUpdateParams(updateParams).
		Where(squirrel.Eq{"id": id}).
		Suffix("returning location_id, photo").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build update animal query")
	}

	commandTag, err := conn.Exec(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "failed to Exec update animal query")
	}

	return commandTag.RowsAffected(), nil
}

func (r *AnimalsRepo) Update(ctx context.Context, id int64, updateParams entity.UpdateAnimal) (int64, error) {
	return r.UpdateWithConn(ctx, r.Pool, id, updateParams)
}

func (r *AnimalsRepo) applyUpdateParams(updateParams entity.UpdateAnimal) squirrel.UpdateBuilder {
	builder := r.Builder.Update(animalsTableName)

	if updateParams.LocationID != nil {
		builder = builder.Set("location_id", *updateParams.LocationID)
	}
	if updateParams.Photo != nil {
		builder = builder.Set("photo", *updateParams.Photo)
	}
	if updateParams.Name != nil {
		builder = builder.Set("name", *updateParams.Name)
	}
	if updateParams.BirthDate != nil {
		builder = builder.Set("birth_date", *updateParams.BirthDate)
	}
	if updateParams.Gender != nil {
		builder = builder.Set("gender", *updateParams.Gender)
	}
	if updateParams.Type != nil {
		builder = builder.Set("type", *updateParams.Type)
	}
	if updateParams.Sterilized != nil {
		builder = builder.Set("sterilized", *updateParams.Sterilized)
	}
	if updateParams.ForAdoption != nil {
		builder = builder.Set("for_adoption", *updateParams.ForAdoption)
	}
	if updateParams.ForWalking != nil {
		builder = builder.Set("for_walking", *updateParams.ForWalking)
	}
	if updateParams.AdopterID != nil {
		builder = builder.Set("adopter_id", *updateParams.AdopterID)
	}
	if updateParams.PublicDescription != nil {
		builder = builder.Set("public_description", *updateParams.PublicDescription)
	}
	if updateParams.PrivateDescription != nil {
		builder = builder.Set("private_description", *updateParams.PrivateDescription)
	}

	return builder
}

func (r *AnimalsRepo) SelectShelterIDWithConn(ctx context.Context, conn usecase.IConnection, animalId int64) (int64, error) {
	sql, args, err := r.Builder.
		Select(fmt.Sprintf("%s.id", sheltersTableName)).
		From(animalsTableName).
		LeftJoin(fmt.Sprintf("%s ON %s.location_id = %s.id", locationsTableName, animalsTableName, locationsTableName)).
		LeftJoin(fmt.Sprintf("%s ON %s.shelter_id = %s.id", sheltersTableName, locationsTableName, sheltersTableName)).
		Where(squirrel.Eq{fmt.Sprintf("%s.id", animalsTableName): animalId}).
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build select animal's shelter_id query")
	}

	var shelterId int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&shelterId)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, errors.Wrap(err, "failed to QueryRow select animal's shelter_id query")
	}

	return shelterId, nil
}

func (r *AnimalsRepo) SelectShelterID(ctx context.Context, animalId int64) (int64, error) {
	return r.SelectShelterIDWithConn(ctx, r.Pool, animalId)
}

func (r *AnimalsRepo) Get(ctx context.Context, id int64) (*entity.Animal, error) {
	sql, args, err := r.Builder.
		Select(fmt.Sprintf("%s.*, %s.shelter_id", animalsTableName, locationsTableName)).
		From(animalsTableName).
		LeftJoin(fmt.Sprintf("%s ON %s.location_id = %s.id", locationsTableName, animalsTableName, locationsTableName)).
		Where(squirrel.Eq{fmt.Sprintf("%s.id", animalsTableName): id}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build get animal query")
	}

	var animal entity.Animal
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&animal.ID, &animal.LocationID, &animal.Photo, &animal.Name,
		&animal.BirthDate, &animal.Type, &animal.Gender, &animal.Sterilized, &animal.ForAdoption, &animal.ForWalking,
		&animal.AdopterID, &animal.PublicDescription, &animal.PrivateDescription, &animal.ShelterID)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to QueryRow get animal query")
	}

	return &animal, nil
}

// DeleteWithConn returns location_id
func (r *AnimalsRepo) DeleteWithConn(ctx context.Context, conn usecase.IConnection, id int64) (locationId int64, err error) {
	sql, args, err := r.Builder.
		Delete(animalsTableName).
		Where(squirrel.Eq{"id": id}).
		Suffix("returning location_id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build delete animal query")
	}

	err = conn.QueryRow(ctx, sql, args...).Scan(&locationId)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return 0, nil
		}

		return 0, errors.Wrap(err, "failed to QueryRow delete animal query")
	}

	return locationId, nil
}
