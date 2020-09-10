package migrater

import (
	"database/sql"

	packr "github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

const pql = "postgres"
const tableName = "public-api"
const migrationPath = "./migrations"

type Migrater struct{}

func (m *Migrater) SetTable() {
	migrate.SetTable(tableName)
}

func (m *Migrater) Source() *migrate.PackrMigrationSource {
	return &migrate.PackrMigrationSource{
		Box: packr.New(tableName, migrationPath),
	}
}

func (m *Migrater) Up(db *sql.DB) (int, error) {
	return migrate.Exec(db, pql, m.Source(), migrate.Up)
}

func (m *Migrater) Down(db *sql.DB) (int, error) {
	return migrate.Exec(db, pql, m.Source(), migrate.Down)
}
