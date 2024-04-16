package usecase

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pets-shelters/backend-svc/pkg/date"
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
		GetLocationsRepo() ILocationsRepo
		GetAnimalsRepo() IAnimalsRepo
		GetAnimalTypesEnumRepo() IAnimalTypesEnumRepo
		GetAdoptersRepo() IAdoptersRepo
		GetTasksRepo() ITasksRepo
		GetTasksAnimalsRepo() ITasksAnimalsRepo
		GetTasksExecutionsRepo() ITasksExecutionsRepo
		Transaction(ctx context.Context, f func(pgx.Tx) error) error
	}

	ISheltersRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, shelter entity.Shelter) (int64, error)
		Create(ctx context.Context, shelter entity.Shelter) (int64, error)
		SelectWithConn(ctx context.Context, conn IConnection) ([]entity.Shelter, error)
		Select(ctx context.Context) ([]entity.Shelter, error)
		Get(ctx context.Context, id int64) (*entity.Shelter, error)
		Update(ctx context.Context, conn IConnection, id int64, updateParams entity.UpdateShelter) (int64, error)
		SelectNames(ctx context.Context, filterName string) ([]entity.SheltersNames, error)
	}

	IUsersRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, user entity.User) (int64, error)
		Create(ctx context.Context, user entity.User) (int64, error)
		Select(ctx context.Context, filters entity.UsersFilters) ([]entity.User, error)
		UpdateShelterIDWithConn(ctx context.Context, conn IConnection, userId int64, shelterId int64) (int64, error)
		Get(ctx context.Context, id int64) (*entity.User, error)
		DeleteWithConn(ctx context.Context, conn IConnection, id int64) (*entity.User, error)
	}

	IFilesRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, file entity.File) (int64, error)
		Create(ctx context.Context, file entity.File) (int64, error)
		Get(ctx context.Context, id int64) (*entity.File, error)
		DeleteWithTemporaryFiles(ctx context.Context, conn IConnection, minCreatedAt time.Time) ([]entity.File, error)
		DeleteWithConn(ctx context.Context, conn IConnection, id int64) error
	}

	ITemporaryFilesRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, temporaryFile entity.TemporaryFile) (int64, error)
		Create(ctx context.Context, temporaryFile entity.TemporaryFile) (int64, error)
		GetWithConn(ctx context.Context, conn IConnection, id int64) (*entity.TemporaryFile, error)
		Get(ctx context.Context, id int64) (*entity.TemporaryFile, error)
		DeleteWithConn(ctx context.Context, conn IConnection, id int64) (*entity.TemporaryFile, error)
		CountForUserId(ctx context.Context, userId int64) (int64, error)
	}

	ILocationsRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, location entity.Location) (int64, error)
		Create(ctx context.Context, location entity.Location) (int64, error)
		GetWithConn(ctx context.Context, conn IConnection, id int64) (*entity.Location, error)
		Get(ctx context.Context, id int64) (*entity.Location, error)
		SelectWithAnimals(ctx context.Context, shelterId int64) ([]entity.LocationsAnimals, error)
		SelectUniqueCities(ctx context.Context) ([]string, error)
		DeleteWithConn(ctx context.Context, conn IConnection, id int64) (*entity.Location, error)
	}

	IAnimalsRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, animal entity.Animal) (int64, error)
		Create(ctx context.Context, animal entity.Animal) (int64, error)
		Select(ctx context.Context, filters entity.AnimalsFilters, pagination *entity.Pagination) ([]entity.AnimalForList, error)
		Count(ctx context.Context, filters entity.AnimalsFilters) (int64, error)
		Update(ctx context.Context, conn IConnection, id int64, updateParams entity.UpdateAnimal) (int64, error)
		SelectShelterID(ctx context.Context, animalId int64) (int64, error)
		Get(ctx context.Context, id int64) (*entity.Animal, error)
		DeleteWithConn(ctx context.Context, conn IConnection, id int64) (locationId int64, err error)
	}

	IAnimalTypesEnumRepo interface {
		Create(ctx context.Context, newValue string) error
		Select(ctx context.Context) ([]string, error)
	}

	IAdoptersRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, adopter entity.Adopter) (int64, error)
		Create(ctx context.Context, adopter entity.Adopter) (int64, error)
		Get(ctx context.Context, id int64) (*entity.Adopter, error)
		Select(ctx context.Context, filterPhoneNumber string) ([]entity.Adopter, error)
	}

	ITasksRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, task entity.Task) (int64, error)
		Create(ctx context.Context, task entity.Task) (int64, error)
		Get(ctx context.Context, id int64) (*entity.Task, error)
		SelectShelterID(ctx context.Context, taskId int64) (int64, error)
		Delete(ctx context.Context, id int64) (int64, error)
		SelectWithExecutions(ctx context.Context, filters entity.TasksFilters) ([]entity.TaskWithExecutions, error)
		SelectForAnimal(ctx context.Context, animalId int64) ([]entity.TaskForAnimal, error)
		SelectForEmails(ctx context.Context, date date.Date) ([]entity.EmployeeTasks, error)
	}

	ITasksAnimalsRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, taskAnimal entity.TaskAnimal) (int64, error)
		Create(ctx context.Context, taskAnimal entity.TaskAnimal) (int64, error)
	}

	ITasksExecutionsRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, taskExecution entity.TaskExecution) (int64, error)
		Create(ctx context.Context, taskExecution entity.TaskExecution) (int64, error)
	}

	IJwt interface {
		CreateTokensPair(userId string) (*structs.TokensPair, error)
		VerifyAccessToken(accessTokenString string) (string, error)
		VerifyRefreshToken(refreshTokenString string) (string, error)
	}

	IAuthorization interface {
		Login() (*structs.LoginResult, error)
		Callback(ctx context.Context, cookie string, googleState string, googleCode string) (*structs.TokensPair, error)
		GetUserInfo(ctx context.Context, userId int64) (*responses.UserInfo, error)
	}

	IShelters interface {
		Create(ctx context.Context, req requests.CreateShelter, userId int64) error
		GetById(ctx context.Context, shelterId int64) (*responses.Shelter, error)
		Update(ctx context.Context, req requests.UpdateShelter, userId int64, shelterId int64) error
		GetNames(ctx context.Context, filterName string) ([]responses.ShelterName, error)
	}

	IEmployees interface {
		Create(ctx context.Context, userId int64, req requests.CreateEmployee) error
		Delete(ctx context.Context, userId int64, idToDelete int64) error
		GetList(ctx context.Context, userId int64) ([]responses.Employee, error)
	}

	ILocations interface {
		Create(ctx context.Context, userId int64, req requests.CreateLocation) error
		GetList(ctx context.Context, shelterId int64) ([]responses.Location, error)
		GetCities(ctx context.Context) ([]string, error)
		Delete(ctx context.Context, userId int64, idToDelete int64) error
	}

	IAnimals interface {
		Create(ctx context.Context, req requests.CreateAnimal, userId int64) error
		GetList(ctx context.Context, filters requests.AnimalsFilters, reqPagination *requests.Pagination) ([]responses.AnimalForList, *responses.PaginationMetadata, error)
		GetTypes(ctx context.Context) (*responses.AnimalTypes, error)
		Update(ctx context.Context, req requests.UpdateAnimal, userId int64, animalId int64) error
		GetById(ctx context.Context, animalId int64, userId *int64) (*responses.Animal, error)
		Delete(ctx context.Context, userId int64, animalId int64) error
	}

	IAdopters interface {
		Create(ctx context.Context, req requests.CreateAdopter) (*responses.AdopterCreated, error)
		GetById(ctx context.Context, adopterId int64) (*responses.Adopter, error)
		GetList(ctx context.Context, filterPhoneNumber string) ([]responses.Adopter, error)
	}

	ITasks interface {
		Create(ctx context.Context, req requests.CreateTask, userId int64) error
		SetExecution(ctx context.Context, req requests.SetTaskDone, taskId int64, userId int64) error
		Delete(ctx context.Context, userId int64, taskId int64) error
		GetListWithExecutions(ctx context.Context, userId int64, req requests.TasksWithExecutionsFilters) ([]responses.TaskWithExecutions, error)
		GetListForAnimal(ctx context.Context, userId int64, animalId int64) ([]responses.TaskForAnimal, error)
	}

	IS3Provider interface {
		UploadFile(ctx context.Context, body io.Reader, bucket string, key string, contentType string) error
		DeleteFile(ctx context.Context, bucket string, key string) error
	}

	IEmailsProvider interface {
		SendInvitationEmail(shelterName string, toEmail string) error
		SendTasksEmail(toEmail string, date date.Date, tasks []structs.TaskForEmail) error
	}

	IFiles interface {
		Upload(ctx context.Context, userId int64, fileContent []byte, fileType string) (int64, error)
	}
)
