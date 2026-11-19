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
       u.created_at,
       ptv.count AS total_count
FROM users u
    INNER JOIN profile_total_views ptv ON u.id = ptv.user_id
WHERE ptv.count > 0
ORDER BY u.id DESC
LIMIT sqlc.arg('limit');

-- name: UsersCreatedAtStatsByDay :many
SELECT g.time                         AS time,
       COALESCE(rcs.count, 0)::BIGINT AS count
FROM (
    SELECT time::TIMESTAMP
    FROM GENERATE_SERIES(
        sqlc.arg('from')::DATE,
        sqlc.arg('to')::DATE,
        '1 DAY'::INTERVAL
    ) AS time
) AS g
    LEFT JOIN (
        SELECT DATE_TRUNC('DAY', created_at) AS time,
               COUNT(*)                      AS count
        FROM users
        WHERE created_at >= sqlc.arg('from')::DATE
        GROUP BY time
    ) AS rcs ON (g.time = rcs.time)
ORDER BY g.time;

-- name: Usernames :many
SELECT u.username
FROM users u
ORDER BY u.id
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: UsersCount :one
SELECT COUNT(*)
FROM users;


-- name: UsersUpsert :one
WITH
    --
    upserted AS (
        INSERT INTO users AS t (social_provider, social_provider_user_id, username, canonical_username, name,
                                created_at, updated_at, last_login_at)
        VALUES (@social_provider, @social_provider_user_id, @username, LOWER(@username::VARCHAR), @name,
                @now::TIMESTAMPTZ, @now::TIMESTAMPTZ, @now::TIMESTAMPTZ)
        ON CONFLICT (social_provider, social_provider_user_id) DO UPDATE
            SET last_login_at      = excluded.last_login_at,
                username           = excluded.username,
                canonical_username = excluded.canonical_username,
                name               = excluded.name,
                updated_at         = CASE
                                         WHEN t.username <> excluded.username OR t.name <> excluded.name
                                             THEN excluded.updated_at
                                         ELSE t.updated_at
                    END
        RETURNING id, (created_at = @now::TIMESTAMPTZ) AS inserted
        --
    ),
    history AS (
        --
        INSERT INTO username_history (user_id, social_provider, canonical_username, created_at, updated_at)
        SELECT id, @social_provider, LOWER(@username::VARCHAR), @now::TIMESTAMPTZ, @now::TIMESTAMPTZ
        FROM upserted
        ON CONFLICT (canonical_username, social_provider, user_id) DO UPDATE
            SET updated_at = excluded.updated_at
        --
    ),
    total_views AS (
        --
        INSERT INTO profile_total_views (user_id, count)
        SELECT id, 0
        FROM upserted
        ON CONFLICT DO NOTHING
        --
    ),
    referral AS (
        --
        INSERT INTO referrals (referee_user_id, referrer_user_id)
        SELECT id, @referrer_user_id::BIGINT
        FROM upserted
        WHERE inserted
          AND @referrer_user_id::BIGINT > 0
          AND EXISTS (SELECT 1 FROM users WHERE id = @referrer_user_id::BIGINT)
        ON CONFLICT DO NOTHING
        --
    )
SELECT id, inserted
FROM upserted;
