package psql

import (
	"fmt"
	"libraries/datastore/migrater"
	"sg/services/public-api/db"
	"sg/services/public-api/models"

	"github.com/Masterminds/squirrel"
	"gopkg.in/gorp.v2"
)

/*
// Possible organization to generalize interfacing with stores
type PsqlConnectionFactory struct {
	APIStoreClient{
		db *gorp.DbMap
		qb *squirrel.StatementBuilderType
	},
	SGClient{},
}
*/

type PsqlConnectionFactory struct {
	db *gorp.DbMap
	qb *squirrel.StatementBuilderType
}

// TODO: Pass in config objects: APIStoreConfig, SGConfig
func NewPsqlConnFactory(host string, conns int) *PsqlConnectionFactory {
	// TODO: Consider adding other clients here
	conn := db.InitDbConn(host, conns)
	qb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// Apply database migrations on Public-API Datastore
	// TODO: Move outside of factory -- does not need to be run every time a new connection is initialized
	m := &migrater.Migrater{}
	count, err := m.Up(conn)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Public-API Datastore: Executed %d migrations\n", count)

	dbMap := &gorp.DbMap{Db: conn, Dialect: gorp.PostgresDialect{}}
	mapTables(dbMap)

	cf := &PsqlConnectionFactory{
		db: dbMap,
		qb: &qb,
	}
	return cf
}

func (f *PsqlConnectionFactory) Connection() *PsqlConnection {
	return &PsqlConnection{
		db: f.db,
		qb: f.qb,
	}
}

// Map golang model to postgres table
// All table primary keys should be id
// TODO: better code reuse?
func mapTables(dbmap *gorp.DbMap) {
	silos := dbmap.AddTableWithName(models.Silo{}, "silos")
	silos.SetKeys(false, "id")
}
