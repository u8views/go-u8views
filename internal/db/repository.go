package db

import (
	"context"
	"database/sql"

	"github.com/u8views/go-u8views/internal/storage/dbs"
)

type Repository struct {
	connection *sql.DB
	queries    *dbs.Queries
}

func NewRepository(connection *sql.DB) (*Repository, error) {
	var queries, err = dbs.Prepare(context.Background(), connection)
	if err != nil {
		return nil, err
	}

	return &Repository{
		connection: connection,
		queries:    queries,
	}, nil
}

func MustRepository(connection *sql.DB) *Repository {
	var repository, err = NewRepository(connection)
	if err != nil {
		panic(err)
	}
	return repository
}

func (r *Repository) Connection() *sql.DB {
	return r.connection
}

func (r *Repository) Queries() *dbs.Queries {
	return r.queries
}

func (r *Repository) Close() error {
	return r.queries.Close()
}

func (r *Repository) WithTransaction(ctx context.Context, fn func(queries *dbs.Queries) error) error {
	return withTransaction(ctx, r.connection, r.queries, fn)
}

func withTransaction(ctx context.Context, db *sql.DB, queries *dbs.Queries, fn func(queries *dbs.Queries) error) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			tx.Rollback()

			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(queries.WithTx(tx))

	return err
}
