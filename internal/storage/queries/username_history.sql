-- name: UsernameHistoryNew :exec
INSERT INTO username_history (user_id, social_provider, canonical_username, created_at, updated_at)
VALUES (@user_id, @social_provider, LOWER(@username), @created_at, @updated_at)
ON CONFLICT (canonical_username, social_provider, user_id) DO UPDATE
    SET updated_at = excluded.updated_at;

-- name: UsernameHistoryGetByOldUsername :one
SELECT u.username
FROM username_history uh
         INNER JOIN users u ON uh.user_id = u.id
WHERE uh.canonical_username = LOWER(@username)
  AND uh.social_provider = @social_provider
ORDER BY uh.updated_at DESC
LIMIT 1;
