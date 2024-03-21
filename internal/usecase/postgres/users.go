package postgres

import (
	"context"
	sql2 "database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/postgres/entity"
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

func (r *UsersRepo) SelectWithConn(ctx context.Context, conn usecase.IConnection) ([]entity.User, error) {
	sql, _, err := r.Builder.
		Select("*").
		From(usersTableName).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build select users query")
	}

	var shelters []entity.User
	rows, err := conn.Query(ctx, sql)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query select users query")
	}

	err = rows.Scan(&shelters)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan users")
	}

	return shelters, nil
}

func (r *UsersRepo) Select(ctx context.Context) ([]entity.User, error) {
	return r.SelectWithConn(ctx, r.Pool)
}

func (r *UsersRepo) UpdateShelterIDWithConn(ctx context.Context, conn usecase.IConnection, userEmail string, shelterId int64) (int64, error) {
	sql, args, err := r.Builder.
		Update(usersTableName).
		Set("shelter_id", shelterId).
		Where(squirrel.Eq{"email": userEmail}).
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
