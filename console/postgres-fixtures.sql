\timing

TRUNCATE profile_hourly_views_stats CASCADE;
TRUNCATE profile_total_views CASCADE;
TRUNCATE users CASCADE;

INSERT INTO users (id)
SELECT generated_id
FROM generate_series(1, 10 * 1000) AS generated_id;

INSERT INTO profile_total_views (user_id, count)
SELECT generated_user_id, 876100
FROM generate_series(1, 10 * 1000) AS generated_user_id;

-- 87,610,000 rows affected in 19 m 8 s 769 ms
INSERT INTO profile_hourly_views_stats (time, user_id, count)
SELECT generated_time, generated_user_id, generated_user_id % 100 + 1
FROM generate_series(
             (DATE_TRUNC('hour', NOW()) - INTERVAL '1 MONTH')::TIMESTAMP,
             (DATE_TRUNC('hour', NOW()))::TIMESTAMP,
             '1 hour'::INTERVAL
         ) AS generated_time
         INNER JOIN
     generate_series(
             1,
             10 * 1000,
             1
         ) AS generated_user_id ON TRUE;

-- 1 MONTH =  7 450 000
-- 1 YEAR  = 87 610 000
SELECT COUNT(*)
FROM profile_hourly_views_stats;
