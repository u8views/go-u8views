package storage

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrations embed.FS

func MigrateUp(db *sql.DB) error {
	// goose has a lot of dependencies with ClickHouse and other DB drivers
	goose.SetBaseFS(migrations)
	goose.SetTableName("schema_migrations")
	// PostgreSQL by default
	// goose.SetDialect("postgres")

	err := goose.Up(db, "migrations")
	if err != nil {
		return err
	}

	return nil
}

func MustMigrateUp(db *sql.DB) {
	var err = MigrateUp(db)

	if err != nil {
		panic(err)
	}
}
