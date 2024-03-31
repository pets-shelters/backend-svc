package usecase

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"io"
	"time"
)

type (
	IConnection interface {
		Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
		Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
		QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
		QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	}

	IDBRepo interface {
		GetSheltersRepo() ISheltersRepo
		GetUsersRepo() IUsersRepo
		GetFilesRepo() IFilesRepo
		GetTemporaryFilesRepo() ITemporaryFilesRepo
		Transaction(ctx context.Context, f func(pgx.Tx) error) error
	}

	ISheltersRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, shelter entity.Shelter) (int64, error)
		Create(ctx context.Context, shelter entity.Shelter) (int64, error)
		SelectWithConn(ctx context.Context, conn IConnection) ([]entity.Shelter, error)
		Select(ctx context.Context) ([]entity.Shelter, error)
	}

	IUsersRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, user entity.User) (int64, error)
		Create(ctx context.Context, user entity.User) (int64, error)
		SelectWithConn(ctx context.Context, conn IConnection, filters entity.UsersFilters) ([]entity.User, error)
		Select(ctx context.Context, filters entity.UsersFilters) ([]entity.User, error)
		UpdateShelterIDWithConn(ctx context.Context, conn IConnection, userId int64, shelterId int64) (int64, error)
	}

	IFilesRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, file entity.File) (int64, error)
		Create(ctx context.Context, file entity.File) (int64, error)
		GetWithConn(ctx context.Context, conn IConnection, id int64) (*entity.File, error)
		Get(ctx context.Context, id int64) (*entity.File, error)
		DeleteWithTemporaryFiles(ctx context.Context, conn IConnection, minCreatedAt time.Time) ([]entity.File, error)
	}

	ITemporaryFilesRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, temporaryFile entity.TemporaryFile) (int64, error)
		Create(ctx context.Context, temporaryFile entity.TemporaryFile) (int64, error)
		GetWithConn(ctx context.Context, conn IConnection, id int64) (*entity.TemporaryFile, error)
		Get(ctx context.Context, id int64) (*entity.TemporaryFile, error)
		DeleteWithConn(ctx context.Context, conn IConnection, id int64) (*entity.TemporaryFile, error)
		CountForUserId(ctx context.Context, userId int64) (int64, error)
	}

	IJwt interface {
		CreateTokensPair(userId string) (*structs.TokensPair, error)
		VerifyAccessToken(accessTokenString string) (string, error)
		VerifyRefreshToken(refreshTokenString string) (string, error)
	}

	IAuthorization interface {
		Login() (*structs.LoginResult, error)
		Callback(ctx context.Context, cookie string, googleState string, googleCode string) (*structs.TokensPair, error)
	}

	IShelters interface {
		Create(ctx context.Context, req requests.CreateShelter, userId int64) error
	}

	IS3Provider interface {
		UploadFile(ctx context.Context, body io.Reader, bucket string, key string, contentType string) error
		DeleteFile(ctx context.Context, bucket string, key string) error
	}

	IFiles interface {
		Upload(ctx context.Context, userId int64, fileContent []byte, fileType string) (int64, error)
	}
)
