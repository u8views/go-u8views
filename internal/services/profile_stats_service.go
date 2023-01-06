package services

import (
	"context"
	"time"

	"github.com/u8views/go-u8views/internal/db"
	"github.com/u8views/go-u8views/internal/storage/dbs"
)

type ProfileStatsCount struct {
	DayCount   int64
	WeekCount  int64
	MonthCount int64
	TotalCount int64
}

type ProfileStatsService struct {
	repository *db.Repository
}

func NewProfileStatsService(repository *db.Repository) *ProfileStatsService {
	return &ProfileStatsService{repository: repository}
}

func (s *ProfileStatsService) StatsCount(ctx context.Context, userID int64, increment bool) (ProfileStatsCount, error) {
	totalCount, err := s.repository.Queries().ProfileTotalViews(ctx, userID)
	if err != nil {
		return ProfileStatsCount{}, err
	}

	now := time.Now().UTC().Truncate(time.Hour)
	stats, err := s.repository.Queries().ProfileHourlyViewsStats(ctx, dbs.ProfileHourlyViewsStatsParams{
		Day:    now.Add(-24 * time.Hour),
		Week:   now.Add(-7 * 24 * time.Hour),
		Month:  now.Add(-31 * 7 * 24 * time.Hour),
		UserID: userID,
	})
	if err != nil {
		return ProfileStatsCount{}, err
	}

	if increment {
		txErr := s.repository.WithTransaction(ctx, func(queries *dbs.Queries) error {
			err := queries.ProfileTotalViewsInc(ctx, userID)
			if err != nil {
				return err
			}

			err = queries.ProfileHourlyViewsStatsUpsert(ctx, dbs.ProfileHourlyViewsStatsUpsertParams{
				Time:   now,
				UserID: userID,
				Count:  1,
			})
			if err != nil {
				return err
			}

			return nil
		})
		if txErr != nil {
			return ProfileStatsCount{}, txErr
		}
	}

	return ProfileStatsCount{
		DayCount:   stats.DayCount,
		WeekCount:  stats.WeekCount,
		MonthCount: stats.MonthCount,
		TotalCount: totalCount,
	}, nil
}
