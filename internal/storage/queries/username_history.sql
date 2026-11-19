-- name: UsernameHistoryGetByOldUsername :one
SELECT u.username
FROM username_history uh
    INNER JOIN users u ON uh.user_id = u.id
WHERE uh.canonical_username = LOWER(@username)
  AND uh.social_provider = @social_provider
ORDER BY uh.updated_at DESC
LIMIT 1;
