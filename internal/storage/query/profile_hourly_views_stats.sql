-- name: ProfileHourlyViewsStatsUpsert :exec
INSERT INTO profile_hourly_views_stats (time, user_id, count)
VALUES (@time, @user_id, @count)
ON CONFLICT (time, user_id) DO UPDATE
    SET count = profile_hourly_views_stats.count + @count;
