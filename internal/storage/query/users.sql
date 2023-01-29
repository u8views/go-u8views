-- name: UsersNew :one
INSERT INTO users (social_provider, social_provider_user_id, username, canonical_username, created_at, updated_at, last_login_at)
VALUES (@social_provider, @social_provider_user_id, @username, LOWER(@username), @created_at, @updated_at, @last_login_at)
ON CONFLICT (social_provider, social_provider_user_id) DO UPDATE
    SET last_login_at = excluded.last_login_at
RETURNING id;

-- name: UsersUpdateUsername :exec
UPDATE users
SET username           = @username,
    canonical_username = LOWER(@username),
    updated_at         = @updated_at
WHERE id = @id
  AND username <> @username;

-- name: UsersGetBySocialProvider :one
SELECT id
FROM users
WHERE social_provider = @social_provider
  AND social_provider_user_id = @social_provider_user_id;

-- name: UsersGetBySocialProviderUsername :one
SELECT id, social_provider_user_id, username
FROM users
WHERE social_provider = @social_provider
  AND canonical_username = LOWER(@username)
ORDER BY last_login_at DESC
LIMIT 1;

