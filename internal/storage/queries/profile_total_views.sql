-- name: ProfileTotalViews :one
SELECT count
FROM profile_total_views
WHERE user_id = @user_id;
