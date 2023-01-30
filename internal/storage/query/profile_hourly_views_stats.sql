-- name: ProfileHourlyViewsStatsUpsert :exec
INSERT INTO profile_hourly_views_stats (user_id, time, count)
VALUES (@user_id, @time, @count)
ON CONFLICT (user_id, time) DO UPDATE
    SET count = profile_hourly_views_stats.count + @count;

-- name: ProfileHourlyViewsStats :many
SELECT user_id,
       SUM(CASE WHEN time >= @day THEN count ELSE 0 END)::BIGINT  AS day_count,
       SUM(CASE WHEN time >= @week THEN count ELSE 0 END)::BIGINT AS week_count,
       SUM(count)::BIGINT                                         AS month_count
FROM profile_hourly_views_stats
WHERE user_id = ANY (@user_ids::BIGINT[])
  AND time >= @month
GROUP BY user_id;
