package repo

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
)

const (
	locationsTableName           = "locations"
	animalLocationConstraintName = "animals_location_id_fkey"
)

type LocationsRepo struct {
	*postgres.Postgres
}

func NewLocationsRepo(pg *postgres.Postgres) *LocationsRepo {
	return &LocationsRepo{pg}
}

func (r *LocationsRepo) CreateWithConn(ctx context.Context, conn usecase.IConnection, location entity.Location) (int64, error) {
	sql, args, err := r.Builder.
		Insert(locationsTableName).
		Columns("city", "address", "shelter_id").
		Values(location.City, location.Address, location.ShelterID).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build location insert query")
	}

	var id int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow location insert query")
	}

	return id, nil
}

func (r *LocationsRepo) Create(ctx context.Context, location entity.Location) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, location)
}

func (r *LocationsRepo) SelectWithAnimals(ctx context.Context, shelterId int64) ([]entity.LocationsAnimals, error) {
	sql, args, err := r.Builder.
		Select(fmt.Sprintf("%s.*, COUNT(%s.*)", locationsTableName, animalsTableName)).
		From(animalsTableName).
		RightJoin(fmt.Sprintf("%s ON %s.location_id = %s.id", locationsTableName, animalsTableName, locationsTableName)).
		Where(squirrel.Eq{fmt.Sprintf("%s.shelter_id", locationsTableName): shelterId}).
		GroupBy(fmt.Sprintf("%s.id", locationsTableName)).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build select locations with animals query")
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query select employees query")
	}

	locationsAnimals := make([]entity.LocationsAnimals, 0)
	defer rows.Close()
	for rows.Next() {
		var locationAnimal entity.LocationsAnimals
		err = rows.Scan(&locationAnimal.ID, &locationAnimal.City, &locationAnimal.Address, &locationAnimal.ShelterID, &locationAnimal.AnimalsNumber)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan locationAnimal entity")
		}

		locationsAnimals = append(locationsAnimals, locationAnimal)
	}

	return locationsAnimals, nil
}

func (r *LocationsRepo) GetWithConn(ctx context.Context, conn usecase.IConnection, id int64) (*entity.Location, error) {
	sql, args, err := r.Builder.
		Select("*").
		From(locationsTableName).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build get location query")
	}

	var location entity.Location
	err = conn.QueryRow(ctx, sql, args...).Scan(&location.ID, &location.City, &location.Address, &location.ShelterID)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to Query get user query")
	}

	return &location, nil
}

func (r *LocationsRepo) Get(ctx context.Context, id int64) (*entity.Location, error) {
	return r.GetWithConn(ctx, r.Pool, id)
}

func (r *LocationsRepo) SelectUniqueCities(ctx context.Context) ([]string, error) {
	sql, args, err := r.Builder.
		Select("DISTINCT ON (city) city").
		From(locationsTableName).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build get locations' unique cities query")
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query select locations' unique cities query")
	}

	cities := make([]string, 0)
	defer rows.Close()
	for rows.Next() {
		var city string
		err = rows.Scan(&city)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan location's city entity")
		}

		cities = append(cities, city)
	}

	return cities, nil
}

func (r *LocationsRepo) DeleteWithConn(ctx context.Context, conn usecase.IConnection, id int64) (*entity.Location, error) {
	sql, args, err := r.Builder.
		Delete(locationsTableName).
		Where(squirrel.Eq{"id": id}).
		Suffix("returning *").
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build delete location query")
	}

	var location entity.Location
	err = conn.QueryRow(ctx, sql, args...).Scan(&location.ID, &location.City, &location.Address, &location.ShelterID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.ConstraintName == animalLocationConstraintName {
			return nil, exceptions.NewLocationHaveAnimalsException()
		}

		if errors.As(err, &pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to QueryRow delete location query")
	}

	return &location, nil
}
