package repo

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
	"time"
)

const filesTableName = "files"

type FilesRepo struct {
	*postgres.Postgres
}

func NewFilesRepo(pg *postgres.Postgres) *FilesRepo {
	return &FilesRepo{pg}
}

func (r *FilesRepo) CreateWithConn(ctx context.Context, conn usecase.IConnection, file entity.File) (int64, error) {
	sql, args, err := r.Builder.
		Insert(filesTableName).
		Columns("bucket", "path").
		Values(file.Bucket, file.Path).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build file insert query")
	}

	var id int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow file insert query")
	}

	return id, nil
}

func (r *FilesRepo) Create(ctx context.Context, file entity.File) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, file)
}

func (r *FilesRepo) GetWithConn(ctx context.Context, conn usecase.IConnection, id int64) (*entity.File, error) {
	sql, args, err := r.Builder.
		Select("*").
		From(filesTableName).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build get file query")
	}

	var file entity.File
	err = conn.QueryRow(ctx, sql, args...).Scan(&file.ID, &file.Bucket, &file.Path)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to QueryRow get file query")
	}

	return &file, nil
}

func (r *FilesRepo) Get(ctx context.Context, id int64) (*entity.File, error) {
	return r.GetWithConn(ctx, r.Pool, id)
}

func (r *FilesRepo) DeleteWithTemporaryFiles(ctx context.Context, conn usecase.IConnection, minCreatedAt time.Time) ([]entity.File, error) {
	sql := fmt.Sprintf("DELETE FROM %s USING %s WHERE %s.id = %s.file_id AND %s.created_at < $1 RETURNING files.*;",
		filesTableName, temporaryFilesTableName, filesTableName, temporaryFilesTableName, temporaryFilesTableName)
	args := []interface{}{minCreatedAt}

	rows, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query delete files query")
	}

	var files []entity.File
	defer rows.Close()
	for rows.Next() {
		var file entity.File
		err = rows.Scan(&file.ID, &file.Bucket, &file.Path)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan file entity")
		}

		files = append(files, file)
	}

	return files, nil
}
