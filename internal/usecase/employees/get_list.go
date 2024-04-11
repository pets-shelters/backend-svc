package employees

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetList(ctx context.Context, userId int64) ([]responses.Employee, error) {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user entity")
	}
	if user == nil {
		return nil, exceptions.NewPermissionDeniedException()
	}

	employees, err := uc.repo.GetUsersRepo().Select(ctx, entity.UsersFilters{
		ShelterID: &user.ShelterID.Int64,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to select users entities")
	}

	return formEmployeesResponse(employees), nil
}

func formEmployeesResponse(employees []entity.User) []responses.Employee {
	response := make([]responses.Employee, 0)
	for _, employee := range employees {
		response = append(response, responses.Employee{
			ID:    employee.ID,
			Email: employee.Email,
			Role:  employee.Role,
		})
	}

	return response
}
