// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: users.sql

package dbs

import (
	"context"
	"time"
)

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
SELECT id, social_provider_user_id, username
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
}

func (q *Queries) UsersGetBySocialProviderUsername(ctx context.Context, arg UsersGetBySocialProviderUsernameParams) (UsersGetBySocialProviderUsernameRow, error) {
	row := q.queryRow(ctx, q.usersGetBySocialProviderUsernameStmt, usersGetBySocialProviderUsername, arg.SocialProvider, arg.Username)
	var i UsersGetBySocialProviderUsernameRow
	err := row.Scan(&i.ID, &i.SocialProviderUserID, &i.Username)
	return i, err
}

const usersNew = `-- name: UsersNew :one
INSERT INTO users (social_provider, social_provider_user_id, username, canonical_username, created_at, updated_at, last_login_at)
VALUES ($1, $2, $3, LOWER($3), $4, $5, $6)
ON CONFLICT (social_provider, social_provider_user_id) DO UPDATE
    SET last_login_at = excluded.last_login_at
RETURNING id
`

type UsersNewParams struct {
	SocialProvider       SocialProvider
	SocialProviderUserID string
	Username             string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	LastLoginAt          time.Time
}

func (q *Queries) UsersNew(ctx context.Context, arg UsersNewParams) (int64, error) {
	row := q.queryRow(ctx, q.usersNewStmt, usersNew,
		arg.SocialProvider,
		arg.SocialProviderUserID,
		arg.Username,
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
    updated_at         = $2
WHERE id = $3
  AND username <> $1
`

type UsersUpdateUsernameParams struct {
	Username  string
	UpdatedAt time.Time
	ID        int64
}

func (q *Queries) UsersUpdateUsername(ctx context.Context, arg UsersUpdateUsernameParams) error {
	_, err := q.exec(ctx, q.usersUpdateUsernameStmt, usersUpdateUsername, arg.Username, arg.UpdatedAt, arg.ID)
	return err
}
