package tasks

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetListWithExecutions(ctx context.Context, userId int64, req requests.TasksWithExecutionsFilters) ([]responses.TaskWithExecutions, error) {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user entity")
	}

	tasksWithExecutions, err := uc.repo.GetTasksRepo().SelectWithExecutions(ctx, entity.TasksFilters{
		ShelterID: &user.ShelterID.Int64,
		Date:      req.Date,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to select tasks entities")
	}

	return formTasksWithExecutionsResponse(tasksWithExecutions), nil
}

func formTasksWithExecutionsResponse(tasksWithExecutions []entity.TaskWithExecutions) []responses.TaskWithExecutions {
	response := make([]responses.TaskWithExecutions, 0)
	for _, taskWithExecution := range tasksWithExecutions {
		responseTaskExecutions := make([]responses.TaskExecution, 0)
		for _, execution := range taskWithExecution.Executions {
			responseTaskExecution := responses.TaskExecution{
				Date:   execution.Date,
				DoneAt: execution.DoneAt,
			}
			if execution.UserID.Valid {
				responseTaskExecution.UserID = &execution.UserID.Int64
			}
			responseTaskExecutions = append(responseTaskExecutions, responseTaskExecution)
		}
		responseTask := responses.TaskWithExecutions{
			ID:          taskWithExecution.ID,
			Description: taskWithExecution.Description,
			AnimalID:    taskWithExecution.AnimalID,
			StartDate:   taskWithExecution.StartDate,
			EndDate:     taskWithExecution.EndDate,
			Time:        taskWithExecution.Time,
			Executions:  responseTaskExecutions,
		}
		response = append(response, responseTask)
	}

	return response
}
