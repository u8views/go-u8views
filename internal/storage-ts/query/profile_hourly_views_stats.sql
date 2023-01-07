-- name: ProfileHourlyViewsStatsUpsert :exec
INSERT INTO profile_hourly_views_stats (user_id, time, count)
VALUES (@user_id, @time, @count)
ON CONFLICT (user_id, time) DO UPDATE
    SET count = profile_hourly_views_stats.count + @count;

-- name: ProfileHourlyViewsStats :one
SELECT COALESCE(SUM(CASE WHEN time >= @day THEN count ELSE 0 END), 0)::BIGINT  AS day_count,
       COALESCE(SUM(CASE WHEN time >= @week THEN count ELSE 0 END), 0)::BIGINT AS week_count,
       COALESCE(SUM(count), 0)::BIGINT                                         AS month_count
FROM profile_hourly_views_stats
WHERE user_id = @user_id
  AND time >= @month;
