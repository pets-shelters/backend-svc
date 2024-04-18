package schedulers

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/date"
	"github.com/pkg/errors"
	"log"
	"time"
)

func (js *JobsScheduler) WithSendWalkRemindersJob(smsProvider usecase.ISmsProvider, atHour int64) error {
	ctx := context.Background()
	_, err := js.scheduler.NewJob(
		gocron.CronJob(
			fmt.Sprintf("0 %d * * *", atHour),
			false,
		),
		gocron.NewTask(
			js.sendWalksReminders(),
			ctx,
			smsProvider,
		),
	)
	if err != nil {
		return errors.Wrap(err, "failed to add new job")
	}

	return nil
}

func (js *JobsScheduler) sendWalksReminders() func(ctx context.Context, smsProvider usecase.ISmsProvider) error {
	return func(ctx context.Context, smsProvider usecase.ISmsProvider) error {
		date := date.Date(time.Now().In(js.location).Add(time.Hour * 24))
		walkings, err := js.repo.GetWalkingsRepo().SelectForReminders(ctx, date)
		if err != nil {
			return errors.Wrap(err, "failed to select walkings for reminders")
		}

		if len(walkings) == 0 {
			return nil
		}

		err = smsProvider.SendWalkReminder(walkings)
		if err != nil {
			log.Printf("semd walks reminders error %+v", err)
			return errors.Wrap(err, "failed to send walks reminders")
		}

		return nil
	}
}
