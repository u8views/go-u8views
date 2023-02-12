-- name: ReferralsNew :exec
INSERT INTO referrals (referee_user_id, referrer_user_id)
VALUES (@referee_user_id, @referrer_user_id)
ON CONFLICT DO NOTHING;

-- name: ReferralsCreatedAtStatsByDay :many
SELECT g.time                        AS time,
       COALESCE(rc.count, 0)::BIGINT AS count
FROM (
    SELECT time::TIMESTAMP
    FROM generate_series(
        sqlc.arg('from')::TIMESTAMP,
        sqlc.arg('to')::TIMESTAMP,
        '1 DAY'::INTERVAL
    ) AS time
) AS g
    LEFT JOIN (
        SELECT DATE_TRUNC('DAY', u.created_at) AS time,
               COUNT(*)                        AS count
        FROM referrals r
                 INNER JOIN users u ON (r.referee_user_id = u.id)
        WHERE r.referrer_user_id = @referrer_user_id
          AND u.created_at >= sqlc.arg('from')::TIMESTAMP
        GROUP BY DATE_TRUNC('DAY', u.created_at)
    ) AS rc ON (g.time = rc.time)
ORDER BY g.time;
