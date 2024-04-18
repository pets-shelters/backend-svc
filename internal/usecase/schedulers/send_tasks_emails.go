package schedulers

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/date"
	"github.com/pkg/errors"
	"time"
)

func (js *JobsScheduler) WithSendTasksEmailsJob(emailsProvider usecase.IEmailsProvider, atHour int64) error {
	ctx := context.Background()
	_, err := js.scheduler.NewJob(
		gocron.CronJob(
			fmt.Sprintf("0 %d * * *", atHour),
			false,
		),
		gocron.NewTask(
			js.sendTasksEmails(),
			ctx,
			emailsProvider,
		),
	)
	if err != nil {
		return errors.Wrap(err, "failed to add new job")
	}

	return nil
}

func (js *JobsScheduler) sendTasksEmails() func(ctx context.Context, emailsProvider usecase.IEmailsProvider) error {
	return func(ctx context.Context, emailsProvider usecase.IEmailsProvider) error {
		date := date.Date(time.Now().In(js.location))
		employeesTasks, err := js.repo.GetTasksRepo().SelectForEmails(ctx, date)
		if err != nil {
			return errors.Wrap(err, "failed to select files for emails")
		}

		for _, employeeTasks := range employeesTasks {
			tasksForEmail := make([]structs.TaskForEmail, 0)
			for _, taskForEmail := range employeeTasks.Tasks {
				structsTaskForEmail := structs.TaskForEmail{
					Description: taskForEmail.Description,
					AnimalName:  taskForEmail.AnimalName,
					AnimalType:  taskForEmail.AnimalType,
					Time:        taskForEmail.Time.String(),
				}
				tasksForEmail = append(tasksForEmail, structsTaskForEmail)
			}

			if len(tasksForEmail) == 0 {
				continue
			}

			err = emailsProvider.SendTasksEmail(employeeTasks.EmployeeEmail, date, tasksForEmail)
			if err != nil {
				return errors.Wrap(err, "failed to send task email")
			}
		}

		return nil
	}
}
