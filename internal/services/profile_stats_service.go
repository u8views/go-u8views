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
	userViewsStatsMap, err := s.UserDayWeekMonthViewsStatsMap(ctx, []int64{userID}, now)
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

	stats := userViewsStatsMap[userID]
	return ProfileViewsStats{
		DayCount:   stats.DayCount,
		WeekCount:  stats.WeekCount,
		MonthCount: stats.MonthCount,
		TotalCount: totalCount,
	}, nil
}

func (s *ProfileStatsService) UserDayWeekMonthViewsStatsMap(ctx context.Context, userIDs []int64, now time.Time) (map[int64]models.DayWeekMonthViewsStats, error) {
	rows, err := s.repository.Queries().ProfileHourlyViewsStats(ctx, dbs.ProfileHourlyViewsStatsParams{
		Day:     now.Add(-24 * time.Hour),
		Week:    now.Add(-7 * 24 * time.Hour),
		Month:   now.Add(-31 * 7 * 24 * time.Hour),
		UserIds: userIDs,
	})
	if err != nil {
		return nil, err
	}

	result := make(map[int64]models.DayWeekMonthViewsStats, len(rows))
	for _, row := range rows {
		result[row.UserID] = models.DayWeekMonthViewsStats{
			DayCount:   row.DayCount,
			WeekCount:  row.WeekCount,
			MonthCount: row.MonthCount,
		}
	}

	return result, nil
}
