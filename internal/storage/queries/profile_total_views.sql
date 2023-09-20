-- name: ProfileTotalViewsNew :exec
INSERT INTO profile_total_views (user_id, count)
VALUES (@user_id, 0)
ON CONFLICT DO NOTHING;

-- name: ProfileTotalViews :one
SELECT count
FROM profile_total_views
WHERE user_id = @user_id;

-- name: ProfileTotalViewsInc :exec
UPDATE profile_total_views
SET count = count + 1
WHERE user_id = @user_id;
