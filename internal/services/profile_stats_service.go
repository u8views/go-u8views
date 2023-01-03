package services

import (
	"context"
	"time"

	"github.com/u8views/go-u8views/internal/db"
	"github.com/u8views/go-u8views/internal/storage/dbs"
)

type ProfileStatsService struct {
	repository *db.Repository
}

func NewProfileStatsService(repository *db.Repository) *ProfileStatsService {
	return &ProfileStatsService{repository: repository}
}

func (s *ProfileStatsService) Count(ctx context.Context, userID int64) (int64, error) {
	count, err := s.repository.Queries().ProfileTotalViews(ctx, userID)
	if err != nil {
		return 0, err
	}

	var now = time.Now().UTC().Truncate(time.Hour)

	txErr := s.repository.WithTransaction(ctx, func(queries *dbs.Queries) error {
		err := queries.ProfileTotalViewsInc(ctx, userID)
		if err != nil {
			return err
		}

		err = queries.ProfileHourlyViewsStatsUpsert(ctx, dbs.ProfileHourlyViewsStatsUpsertParams{
			Time:   now,
			UserID: userID,
			Count:  count,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return 0, txErr
	}

	return count, nil
}
