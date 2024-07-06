-- name: ProfileTimePeriodViews :one
SELECT COALESCE(SUM(count), 0)::BIGINT as count
FROM profile_hourly_views_stats
WHERE user_id = @user_id
  AND time >= @time_period;