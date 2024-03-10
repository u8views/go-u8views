-- name: UserNicknamesNew :exec
INSERT INTO user_nicknames (user_id, old_nickname)
VALUES (@user_id, @old_nickname)
ON CONFLICT DO NOTHING;

-- name: UserNicknamesGetByOldUsername :one
SELECT u.id, u.social_provider_user_id, u.username, u.name
FROM user_nicknames 
JOIN users u
    ON user_id = u.id
WHERE old_nickname = @old_nickname
    AND u.social_provider = @social_provider
ORDER BY u.last_login_at DESC
LIMIT 1;

-- name: UserNicknamesDelete :exec
DELETE FROM user_nicknames
WHERE old_nickname = @old_nickname
  AND user_id IN (
    SELECT id FROM users
    WHERE social_provider = @social_provider
  );

