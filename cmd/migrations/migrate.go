package migrations

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var embedMigrations embed.FS

func Migrate(conn *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("clickhouse"); err != nil {
		return err
	}

	if err := goose.Up(conn, ".", goose.WithNoVersioning()); err != nil {
		return err
	}

	return nil
}
