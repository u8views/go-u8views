TRUNCATE profile_total_views CASCADE;
TRUNCATE users CASCADE;

INSERT INTO users (id)
SELECT *
FROM generate_series(1, 10 * 1000);

INSERT INTO profile_total_views (user_id, count)
SELECT *, 0
FROM generate_series(1, 10 * 1000);
