package repo

import (
	"context"
	sql2 "database/sql"
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
	usersTableName        = "users"
	emailUniqueConstraint = "unique_user_email"
)

type UsersRepo struct {
	*postgres.Postgres
}

func NewUsersRepo(pg *postgres.Postgres) *UsersRepo {
	return &UsersRepo{pg}
}

func (r *UsersRepo) CreateWithConn(ctx context.Context, conn usecase.IConnection, user entity.User) (int64, error) {
	sql, args, err := r.Builder.
		Insert(usersTableName).
		Columns("email", "role", "shelter_id").
		Values(user.Email, user.Role, user.ShelterID).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build user insert query")
	}

	var id int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.ConstraintName == emailUniqueConstraint {
			return 0, exceptions.NewUserExistsException()
		}
		return 0, errors.Wrap(err, "failed to QueryRow user insert query")
	}

	return id, nil
}

func (r *UsersRepo) Create(ctx context.Context, user entity.User) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, user)
}

func (r *UsersRepo) SelectWithConn(ctx context.Context, conn usecase.IConnection, filters entity.UsersFilters) ([]entity.User, error) {
	sql, args, err := r.applyFilters(filters).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build select users query")
	}

	rows, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query select users query")
	}

	var users []entity.User
	defer rows.Close()
	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.Email, &user.ShelterID, &user.Role)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan user entity")
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UsersRepo) Select(ctx context.Context, filters entity.UsersFilters) ([]entity.User, error) {
	return r.SelectWithConn(ctx, r.Pool, filters)
}

func (r *UsersRepo) applyFilters(filters entity.UsersFilters) squirrel.SelectBuilder {
	builder := r.Builder.Select("*").From(usersTableName)
	if filters.Email != nil {
		builder = builder.Where(squirrel.Eq{"email": *filters.Email})
	}
	if filters.ShelterID != nil {
		builder = builder.Where(squirrel.Eq{"shelter_id": *filters.ShelterID})
	}

	return builder
}

func (r *UsersRepo) Get(ctx context.Context, conn usecase.IConnection, id int64) (*entity.User, error) {
	sql, args, err := r.Builder.
		Select("*").
		From(usersTableName).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build get user query")
	}

	var user entity.User
	err = conn.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.Email, &user.ShelterID, &user.Role)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to Query get user query")
	}

	return &user, nil
}

func (r *UsersRepo) UpdateShelterIDWithConn(ctx context.Context, conn usecase.IConnection, userId int64, shelterId int64) (int64, error) {
	sql, args, err := r.Builder.
		Update(usersTableName).
		Set("shelter_id", shelterId).
		Where(squirrel.Eq{"id": userId}).
		Where(squirrel.Eq{"shelter_id": sql2.NullInt64{}}).
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build user update shelter_id query")
	}

	commandTag, err := conn.Exec(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "failed to Exec user update shelter_id query")
	}

	return commandTag.RowsAffected(), nil
}
