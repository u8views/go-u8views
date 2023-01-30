-- name: UsersNew :one
INSERT INTO users (social_provider, social_provider_user_id, username, canonical_username, name, created_at, updated_at,
                   last_login_at)
VALUES (@social_provider, @social_provider_user_id, @username, LOWER(@username::VARCHAR), @name, @created_at,
        @updated_at, @last_login_at)
ON CONFLICT (social_provider, social_provider_user_id) DO UPDATE
    SET last_login_at = excluded.last_login_at
RETURNING id;

-- name: UsersUpdateUsername :exec
UPDATE users
SET username           = @username,
    canonical_username = LOWER(@username),
    name               = @name,
    updated_at         = @updated_at
WHERE id = @id
  AND (username <> @username OR name <> @name);

-- name: UsersGetBySocialProvider :one
SELECT id
FROM users
WHERE social_provider = @social_provider
  AND social_provider_user_id = @social_provider_user_id;

-- name: UsersGetByID :one
SELECT id, social_provider_user_id, username, name
FROM users
WHERE id = @id;

-- name: UsersGetBySocialProviderUsername :one
SELECT id, social_provider_user_id, username, name
FROM users
WHERE social_provider = @social_provider
  AND canonical_username = LOWER(@username)
ORDER BY last_login_at DESC
LIMIT 1;

-- name: UsersGet :many
SELECT u.id,
       u.social_provider_user_id,
       u.username,
       u.name,
       ptv.count
FROM users u
         INNER JOIN profile_total_views ptv ON u.id = ptv.user_id
WHERE ptv.count > 0
ORDER BY u.id DESC
LIMIT 32;
