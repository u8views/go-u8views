package services

import (
	"context"
	"time"

	"github.com/u8views/go-u8views/internal/db"
	"github.com/u8views/go-u8views/internal/models"
	"github.com/u8views/go-u8views/internal/storage/dbs"
)

type ProfileViewsStats = models.ProfileViewsStats

type ProfileStatsService struct {
	repository *db.Repository
}

func NewProfileStatsService(repository *db.Repository) *ProfileStatsService {
	return &ProfileStatsService{repository: repository}
}

func (s *ProfileStatsService) StatsCount(ctx context.Context, userID int64, increment bool) (result ProfileViewsStats, err error) {
	totalCount, err := s.repository.Queries().ProfileTotalViews(ctx, userID)
	if err != nil {
		return result, err
	}

	now := time.Now().UTC().Truncate(time.Hour)
	stats, err := s.repository.Queries().ProfileHourlyViewsStats(ctx, dbs.ProfileHourlyViewsStatsParams{
		Day:    now.Add(-24 * time.Hour),
		Week:   now.Add(-7 * 24 * time.Hour),
		Month:  now.Add(-31 * 7 * 24 * time.Hour),
		UserID: userID,
	})
	if err != nil {
		return result, err
	}

	if increment {
		txErr := s.repository.WithTransaction(ctx, func(queries *dbs.Queries) error {
			err := queries.ProfileTotalViewsInc(ctx, userID)
			if err != nil {
				return err
			}

			err = queries.ProfileHourlyViewsStatsUpsert(ctx, dbs.ProfileHourlyViewsStatsUpsertParams{
				UserID: userID,
				Time:   now,
				Count:  1,
			})
			if err != nil {
				return err
			}

			return nil
		})
		if txErr != nil {
			return result, txErr
		}
	}

	return ProfileViewsStats{
		DayCount:   stats.DayCount,
		WeekCount:  stats.WeekCount,
		MonthCount: stats.MonthCount,
		TotalCount: totalCount,
	}, nil
}
