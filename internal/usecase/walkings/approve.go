package walkings

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
)

func (uc *UseCase) Approve(ctx context.Context, req requests.ApproveWalking, userId int64, walkingId int64) error {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to get user entity")
	}
	if user == nil {
		return exceptions.NewPermissionDeniedException()
	}

	approvedStatus := structs.ApprovedWalkingStatus
	err = uc.repo.Transaction(ctx, func(tx pgx.Tx) error {
		animalId, err := uc.repo.GetWalkingsRepo().Update(ctx, tx, walkingId, entity.UpdateWalking{
			Status: &approvedStatus,
			Date:   req.Date,
			Time:   &req.Time,
		})
		if err != nil {
			return errors.Wrap(err, "failed to update walking entity")
		}
		if animalId == 0 {
			return exceptions.NewNotFoundException()
		}

		shelterId, err := uc.repo.GetAnimalsRepo().SelectShelterIDWithConn(ctx, tx, animalId)
		if err != nil {
			return errors.Wrap(err, "failed to select animal's shelter_id")
		}
		if user.ShelterID.Int64 != shelterId {
			return exceptions.NewPermissionDeniedException()
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to process transaction")
	}

	return nil
}
