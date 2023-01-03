package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connection(dsn string) (*sql.DB, error) {
	return sql.Open("postgres", dsn)
}

func MustConnection(dsn string) *sql.DB {
	var connection, err = Connection(dsn)

	if err != nil {
		panic(err)
	}

	return connection
}
