\timing

-- 10 000
SELECT 'users' AS table_name, COUNT(*)
FROM users;

-- 10 000
SELECT 'profile_total_views' AS table_name, COUNT(*)
FROM profile_total_views;

-- 87 610 000
SELECT 'profile_hourly_views_stats' AS table_name, COUNT(*)
FROM profile_hourly_views_stats;
