package schedulers

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pkg/errors"
	"time"
)

type JobsScheduler struct {
	scheduler gocron.Scheduler
	repo      usecase.IDBRepo
}

func NewJobsScheduler(logger gocron.Logger, repo usecase.IDBRepo, location *time.Location) (*JobsScheduler, error) {
	runner, err := gocron.NewScheduler(
		gocron.WithLogger(
			logger,
		),
		gocron.WithLocation(
			location,
		),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new scheduler")
	}

	return &JobsScheduler{
		scheduler: runner,
		repo:      repo,
	}, nil
}

func (js *JobsScheduler) Start() {
	js.scheduler.Start()
}

func (js *JobsScheduler) Shutdown() {
	js.scheduler.Shutdown()
}
