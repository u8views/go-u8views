// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package dbs

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.profileHourlyViewsStatsStmt, err = db.PrepareContext(ctx, profileHourlyViewsStats); err != nil {
		return nil, fmt.Errorf("error preparing query ProfileHourlyViewsStats: %w", err)
	}
	if q.profileHourlyViewsStatsUpsertStmt, err = db.PrepareContext(ctx, profileHourlyViewsStatsUpsert); err != nil {
		return nil, fmt.Errorf("error preparing query ProfileHourlyViewsStatsUpsert: %w", err)
	}
	if q.profileTotalViewsStmt, err = db.PrepareContext(ctx, profileTotalViews); err != nil {
		return nil, fmt.Errorf("error preparing query ProfileTotalViews: %w", err)
	}
	if q.profileTotalViewsIncStmt, err = db.PrepareContext(ctx, profileTotalViewsInc); err != nil {
		return nil, fmt.Errorf("error preparing query ProfileTotalViewsInc: %w", err)
	}
	if q.profileTotalViewsNewStmt, err = db.PrepareContext(ctx, profileTotalViewsNew); err != nil {
		return nil, fmt.Errorf("error preparing query ProfileTotalViewsNew: %w", err)
	}
	if q.usersGetBySocialProviderStmt, err = db.PrepareContext(ctx, usersGetBySocialProvider); err != nil {
		return nil, fmt.Errorf("error preparing query UsersGetBySocialProvider: %w", err)
	}
	if q.usersGetBySocialProviderUsernameStmt, err = db.PrepareContext(ctx, usersGetBySocialProviderUsername); err != nil {
		return nil, fmt.Errorf("error preparing query UsersGetBySocialProviderUsername: %w", err)
	}
	if q.usersNewStmt, err = db.PrepareContext(ctx, usersNew); err != nil {
		return nil, fmt.Errorf("error preparing query UsersNew: %w", err)
	}
	if q.usersUpdateUsernameStmt, err = db.PrepareContext(ctx, usersUpdateUsername); err != nil {
		return nil, fmt.Errorf("error preparing query UsersUpdateUsername: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.profileHourlyViewsStatsStmt != nil {
		if cerr := q.profileHourlyViewsStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing profileHourlyViewsStatsStmt: %w", cerr)
		}
	}
	if q.profileHourlyViewsStatsUpsertStmt != nil {
		if cerr := q.profileHourlyViewsStatsUpsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing profileHourlyViewsStatsUpsertStmt: %w", cerr)
		}
	}
	if q.profileTotalViewsStmt != nil {
		if cerr := q.profileTotalViewsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing profileTotalViewsStmt: %w", cerr)
		}
	}
	if q.profileTotalViewsIncStmt != nil {
		if cerr := q.profileTotalViewsIncStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing profileTotalViewsIncStmt: %w", cerr)
		}
	}
	if q.profileTotalViewsNewStmt != nil {
		if cerr := q.profileTotalViewsNewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing profileTotalViewsNewStmt: %w", cerr)
		}
	}
	if q.usersGetBySocialProviderStmt != nil {
		if cerr := q.usersGetBySocialProviderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing usersGetBySocialProviderStmt: %w", cerr)
		}
	}
	if q.usersGetBySocialProviderUsernameStmt != nil {
		if cerr := q.usersGetBySocialProviderUsernameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing usersGetBySocialProviderUsernameStmt: %w", cerr)
		}
	}
	if q.usersNewStmt != nil {
		if cerr := q.usersNewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing usersNewStmt: %w", cerr)
		}
	}
	if q.usersUpdateUsernameStmt != nil {
		if cerr := q.usersUpdateUsernameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing usersUpdateUsernameStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                   DBTX
	tx                                   *sql.Tx
	profileHourlyViewsStatsStmt          *sql.Stmt
	profileHourlyViewsStatsUpsertStmt    *sql.Stmt
	profileTotalViewsStmt                *sql.Stmt
	profileTotalViewsIncStmt             *sql.Stmt
	profileTotalViewsNewStmt             *sql.Stmt
	usersGetBySocialProviderStmt         *sql.Stmt
	usersGetBySocialProviderUsernameStmt *sql.Stmt
	usersNewStmt                         *sql.Stmt
	usersUpdateUsernameStmt              *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                   tx,
		tx:                                   tx,
		profileHourlyViewsStatsStmt:          q.profileHourlyViewsStatsStmt,
		profileHourlyViewsStatsUpsertStmt:    q.profileHourlyViewsStatsUpsertStmt,
		profileTotalViewsStmt:                q.profileTotalViewsStmt,
		profileTotalViewsIncStmt:             q.profileTotalViewsIncStmt,
		profileTotalViewsNewStmt:             q.profileTotalViewsNewStmt,
		usersGetBySocialProviderStmt:         q.usersGetBySocialProviderStmt,
		usersGetBySocialProviderUsernameStmt: q.usersGetBySocialProviderUsernameStmt,
		usersNewStmt:                         q.usersNewStmt,
		usersUpdateUsernameStmt:              q.usersUpdateUsernameStmt,
	}
}
