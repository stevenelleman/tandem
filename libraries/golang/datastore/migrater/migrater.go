package migrater

import (
	"database/sql"

	packr "github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

type Args struct {
	Dialect   string
	TableName string
	Path      string
}

func MakeArgs(dialect, tableName, path string) *Args {
	return &Args{
		Dialect:   dialect,
		TableName: tableName,
		Path:      path,
	}
}

type Migrater struct {
	dialect   string
	tableName string
	path      string // path to migrations
}

func MakeMigrater(args *Args) *Migrater {
	return &Migrater{
		dialect:   args.Dialect,
		tableName: args.TableName,
		path:      args.Path,
	}
}

func (m *Migrater) SetTable() {
	migrate.SetTable(m.tableName)
}

func (m *Migrater) Source() *migrate.PackrMigrationSource {
	return &migrate.PackrMigrationSource{
		Box: packr.New(m.tableName, m.path),
	}
}

func (m *Migrater) Up(db *sql.DB) (int, error) {
	return migrate.Exec(db, m.dialect, m.Source(), migrate.Up)
}

func (m *Migrater) Down(db *sql.DB) (int, error) {
	return migrate.Exec(db, m.dialect, m.Source(), migrate.Down)
}
