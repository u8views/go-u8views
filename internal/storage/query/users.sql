-- name: UsersNew :one
INSERT INTO users (social_provider, social_provider_user_id, username, created_at, updated_at, last_login_at)
VALUES (@social_provider, @social_provider_user_id, @username, @created_at, @updated_at, @last_login_at)
ON CONFLICT (social_provider, social_provider_user_id) DO UPDATE
    SET last_login_at = excluded.last_login_at
RETURNING id;

-- name: UsersUpdateUsername :exec
UPDATE users
SET username   = @username,
    updated_at = @updated_at
WHERE id = @id
  AND username <> @username;
