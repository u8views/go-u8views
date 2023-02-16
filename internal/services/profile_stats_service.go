package services

import (
	"context"
	"log"
	"time"

	"github.com/u8views/go-u8views/internal/db"
	"github.com/u8views/go-u8views/internal/models"
	"github.com/u8views/go-u8views/internal/storage/dbs"
)

type ProfileViewsStats = models.ProfileViewsStats

type StatsService struct {
	repository *db.Repository
}

func NewStatsService(repository *db.Repository) *StatsService {
	return &StatsService{repository: repository}
}

func (s *StatsService) StatsCount(ctx context.Context, userID int64, increment bool) (result ProfileViewsStats, err error) {
	totalCount, err := s.repository.Queries().ProfileTotalViews(ctx, userID)
	if err != nil {
		return result, err
	}

	now := time.Now().UTC().Truncate(time.Hour)
	userViewsStatsMap, err := s.UserDayWeekMonthViewsStatsMap(ctx, []int64{userID}, now)
	if err != nil {
		return result, err
	}

	stats := userViewsStatsMap[userID]

	if increment {
		stats.Inc()
		totalCount += 1

		go func() {
			err := s.increment(context.Background(), userID, now)
			if err != nil {
				log.Printf("Database err %s\n", err)
			}
		}()
	}

	return ProfileViewsStats{
		DayCount:   stats.DayCount,
		WeekCount:  stats.WeekCount,
		MonthCount: stats.MonthCount,
		TotalCount: totalCount,
	}, nil
}

func (s *StatsService) UserDayWeekMonthViewsStatsMap(ctx context.Context, userIDs []int64, now time.Time) (map[int64]models.DayWeekMonthViewsStats, error) {
	rows, err := s.repository.Queries().ProfileHourlyViewsStats(ctx, dbs.ProfileHourlyViewsStatsParams{
		Day:     now.AddDate(0, 0, -1),
		Week:    now.AddDate(0, 0, -7),
		Month:   now.AddDate(0, -1, 0),
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

func (s *StatsService) Stats(ctx context.Context, userID int64) ([]models.TimeCount, error) {
	to := time.Now().UTC().Truncate(time.Hour)
	from := to.AddDate(0, -1, 0)

	rows, err := s.repository.Queries().ProfileHourlyViewsStatsByHour(ctx, dbs.ProfileHourlyViewsStatsByHourParams{
		UserID: userID,
		From:   from,
		To:     to,
	})
	if err != nil {
		return nil, err
	}

	result := make([]models.TimeCount, len(rows))
	for i, row := range rows {
		result[i] = models.TimeCount{
			Time:  row.Time.Unix(),
			Count: row.Count,
		}
	}
	return result, nil
}

func (s *StatsService) ReferralsStats(ctx context.Context, userID int64) ([]models.TimeCount, error) {
	to := time.Now().UTC().Truncate(24 * time.Hour)
	from := to.AddDate(0, -1, 0)

	rows, err := s.repository.Queries().ReferralsCreatedAtStatsByDay(ctx, dbs.ReferralsCreatedAtStatsByDayParams{
		ReferrerUserID: userID,
		From:           from,
		To:             to,
	})
	if err != nil {
		return nil, err
	}

	result := make([]models.TimeCount, len(rows))
	for i, row := range rows {
		result[i] = models.TimeCount{
			Time:  row.Time.Unix(),
			Count: row.Count,
		}
	}
	return result, nil
}

func (s *StatsService) increment(ctx context.Context, userID int64, now time.Time) error {
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
		return txErr
	}

	return nil
}
