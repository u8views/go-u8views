-- name: ProfileHourlyViewsStatsUpsert :exec
INSERT INTO profile_hourly_views_stats (time, user_id, count)
VALUES (@time, @user_id, @count)
ON CONFLICT (time, user_id) DO UPDATE
    SET count = profile_hourly_views_stats.count + @count;

-- name: ProfileHourlyViewsStats :one
SELECT COALESCE(SUM(CASE WHEN time >= @day THEN count ELSE 0 END), 0)::BIGINT  AS day_count,
       COALESCE(SUM(CASE WHEN time >= @week THEN count ELSE 0 END), 0)::BIGINT AS week_count,
       COALESCE(SUM(count), 0)::BIGINT                                         AS month_count
FROM profile_hourly_views_stats
WHERE time >= @month
  AND user_id = @user_id;
