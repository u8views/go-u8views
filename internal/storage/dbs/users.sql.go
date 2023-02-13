// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: users.sql

package dbs

import (
	"context"
	"time"
)

const usersCreatedAtStatsByDay = `-- name: UsersCreatedAtStatsByDay :many
SELECT g.time                         AS time,
       COALESCE(rcs.count, 0)::BIGINT AS count
FROM (
    SELECT time::TIMESTAMP
    FROM generate_series(
        $1::TIMESTAMP,
        $2::TIMESTAMP,
        '1 DAY'::INTERVAL
    ) AS time
) AS g
    LEFT JOIN (
        SELECT DATE_TRUNC('DAY', created_at) AS time,
               COUNT(*)                      AS count
        FROM users
        WHERE created_at >= $1::TIMESTAMP
        GROUP BY time
    ) AS rcs ON (g.time = rcs.time)
ORDER BY g.time
`

type UsersCreatedAtStatsByDayParams struct {
	From time.Time
	To   time.Time
}

type UsersCreatedAtStatsByDayRow struct {
	Time  time.Time
	Count int64
}

func (q *Queries) UsersCreatedAtStatsByDay(ctx context.Context, arg UsersCreatedAtStatsByDayParams) ([]UsersCreatedAtStatsByDayRow, error) {
	rows, err := q.query(ctx, q.usersCreatedAtStatsByDayStmt, usersCreatedAtStatsByDay, arg.From, arg.To)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UsersCreatedAtStatsByDayRow
	for rows.Next() {
		var i UsersCreatedAtStatsByDayRow
		if err := rows.Scan(&i.Time, &i.Count); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const usersGet = `-- name: UsersGet :many
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
LIMIT $1
`

type UsersGetRow struct {
	ID                   int64
	SocialProviderUserID string
	Username             string
	Name                 string
	CreatedAt            time.Time
	TotalCount           int64
}

func (q *Queries) UsersGet(ctx context.Context, limit int32) ([]UsersGetRow, error) {
	rows, err := q.query(ctx, q.usersGetStmt, usersGet, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UsersGetRow
	for rows.Next() {
		var i UsersGetRow
		if err := rows.Scan(
			&i.ID,
			&i.SocialProviderUserID,
			&i.Username,
			&i.Name,
			&i.CreatedAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UsersGetAllUsernames = `-- name: UsersGetAllUsernames :many
SELECT u.username FROM users u 
ORDER BY u.id DESC;
`

func (q *Queries) GetAllUsernames(ctx context.Context) ([]UsersGetRow, error) {
	rows, err := q.query(ctx, q.usersGetAllUsernamesStmt, UsersGetAllUsernames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UsersGetRow
	for rows.Next() {
		var i UsersGetRow
		if err := rows.Scan(
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const usersGetByID = `-- name: UsersGetByID :one
SELECT id, social_provider_user_id, username, name
FROM users
WHERE id = $1
`

type UsersGetByIDRow struct {
	ID                   int64
	SocialProviderUserID string
	Username             string
	Name                 string
}

func (q *Queries) UsersGetByID(ctx context.Context, id int64) (UsersGetByIDRow, error) {
	row := q.queryRow(ctx, q.usersGetByIDStmt, usersGetByID, id)
	var i UsersGetByIDRow
	err := row.Scan(
		&i.ID,
		&i.SocialProviderUserID,
		&i.Username,
		&i.Name,
	)
	return i, err
}

const usersGetBySocialProvider = `-- name: UsersGetBySocialProvider :one
SELECT id
FROM users
WHERE social_provider = $1
  AND social_provider_user_id = $2
`

type UsersGetBySocialProviderParams struct {
	SocialProvider       SocialProvider
	SocialProviderUserID string
}

func (q *Queries) UsersGetBySocialProvider(ctx context.Context, arg UsersGetBySocialProviderParams) (int64, error) {
	row := q.queryRow(ctx, q.usersGetBySocialProviderStmt, usersGetBySocialProvider, arg.SocialProvider, arg.SocialProviderUserID)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const usersGetBySocialProviderUsername = `-- name: UsersGetBySocialProviderUsername :one
SELECT id, social_provider_user_id, username, name
FROM users
WHERE social_provider = $1
  AND canonical_username = LOWER($2)
ORDER BY last_login_at DESC
LIMIT 1
`

type UsersGetBySocialProviderUsernameParams struct {
	SocialProvider SocialProvider
	Username       string
}

type UsersGetBySocialProviderUsernameRow struct {
	ID                   int64
	SocialProviderUserID string
	Username             string
	Name                 string
}

func (q *Queries) UsersGetBySocialProviderUsername(ctx context.Context, arg UsersGetBySocialProviderUsernameParams) (UsersGetBySocialProviderUsernameRow, error) {
	row := q.queryRow(ctx, q.usersGetBySocialProviderUsernameStmt, usersGetBySocialProviderUsername, arg.SocialProvider, arg.Username)
	var i UsersGetBySocialProviderUsernameRow
	err := row.Scan(
		&i.ID,
		&i.SocialProviderUserID,
		&i.Username,
		&i.Name,
	)
	return i, err
}

const usersNew = `-- name: UsersNew :one
INSERT INTO users (social_provider, social_provider_user_id, username, canonical_username, name, created_at, updated_at,
                   last_login_at)
VALUES ($1, $2, $3, LOWER($3::VARCHAR), $4, $5,
        $6, $7)
ON CONFLICT (social_provider, social_provider_user_id) DO UPDATE
    SET last_login_at = excluded.last_login_at
RETURNING id
`

type UsersNewParams struct {
	SocialProvider       SocialProvider
	SocialProviderUserID string
	Username             string
	Name                 string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	LastLoginAt          time.Time
}

func (q *Queries) UsersNew(ctx context.Context, arg UsersNewParams) (int64, error) {
	row := q.queryRow(ctx, q.usersNewStmt, usersNew,
		arg.SocialProvider,
		arg.SocialProviderUserID,
		arg.Username,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.LastLoginAt,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const usersUpdateUsername = `-- name: UsersUpdateUsername :exec
UPDATE users
SET username           = $1,
    canonical_username = LOWER($1),
    name               = $2,
    updated_at         = $3
WHERE id = $4
  AND (username <> $1 OR name <> $2)
`

type UsersUpdateUsernameParams struct {
	Username  string
	Name      string
	UpdatedAt time.Time
	ID        int64
}

func (q *Queries) UsersUpdateUsername(ctx context.Context, arg UsersUpdateUsernameParams) error {
	_, err := q.exec(ctx, q.usersUpdateUsernameStmt, usersUpdateUsername,
		arg.Username,
		arg.Name,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
