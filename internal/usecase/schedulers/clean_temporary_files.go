package schedulers

import (
	"context"
	"github.com/go-co-op/gocron/v2"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/configs"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pkg/errors"
	"log"
	"time"
)

func (js *JobsScheduler) WithCleanTemporaryFilesJob(s3Provider usecase.IS3Provider, cfg configs.TemporaryFiles) error {
	ctx := context.Background()
	_, err := js.scheduler.NewJob(
		gocron.DurationJob(
			cfg.SchedulerPeriod,
		),
		gocron.NewTask(
			js.cleanTemporaryFiles(),
			ctx,
			s3Provider,
			cfg.Lifetime,
		),
	)
	if err != nil {
		return errors.Wrap(err, "failed to add new job")
	}

	return nil
}

func (js *JobsScheduler) cleanTemporaryFiles() func(ctx context.Context, s3Provider usecase.IS3Provider, temporaryFileLifetime time.Duration) {
	return func(ctx context.Context, s3Provider usecase.IS3Provider, temporaryFileLifetime time.Duration) {
		minCreatedAt := time.Now().UTC().Add(-temporaryFileLifetime)
		js.repo.Transaction(ctx, func(tx pgx.Tx) error {
			files, err := js.repo.GetFilesRepo().DeleteWithTemporaryFiles(ctx, tx, minCreatedAt)
			if err != nil {
				log.Print(err)
				return errors.Wrap(err, "failed to delete files with temporary_files")
			}

			for _, file := range files {
				err = s3Provider.DeleteFile(ctx, file.Bucket, file.Path)
				if err != nil {
					return errors.Wrap(err, "failed to delete file from s3")
				}
			}

			return nil
		})
	}
}
