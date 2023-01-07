package tests

import (
	"context"
	"sync/atomic"
	"testing"

	"github.com/u8views/go-u8views/internal/db"
	"github.com/u8views/go-u8views/internal/env"
	"github.com/u8views/go-u8views/internal/services"

	"github.com/stretchr/testify/require"
)

func BenchmarkProfileStatsService(b *testing.B) {
	var (
		dsn = env.Must("DSN")
	)

	var pgConnection = db.MustConnection(dsn)
	defer pgConnection.Close()

	var repository = db.MustRepository(pgConnection)
	defer repository.Close()

	var service = services.NewProfileStatsService(repository)

	var count int64

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var userID = atomic.AddInt64(&count, 1)%10000 + 1

			stats, err := service.StatsCount(context.Background(), userID, false)

			require.NoError(b, err)
			require.True(b, stats.DayCount > 0)
			require.True(b, stats.WeekCount > 0)
			require.True(b, stats.MonthCount > 0)
			require.True(b, stats.TotalCount > 0)
		}
	})
}
