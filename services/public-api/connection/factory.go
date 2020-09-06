package connection

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"gopkg.in/gorp.v2"
	migrater "libraries/datastore/migrater"
	"sg/services/public-api/db"
	"sg/services/public-api/models"
)

type ConnectionFactory struct {
	db *gorp.DbMap
	qb *squirrel.StatementBuilderType
	mg *migrater.Migrater
}

func NewConnectionFactory(host string, conns int) *ConnectionFactory {
	database := db.InitDb(host, conns)
	qb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	m := &migrater.Migrater{}

	// Apply database migrations on Public-API Datastore
	count, err := m.Up(database)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Public-API Datastore: Executed %d migrations\n", count)

	dbMap := &gorp.DbMap{Db: database, Dialect: gorp.PostgresDialect{}}
	mapTables(dbMap)

	cf := &ConnectionFactory{
		db: dbMap,
		qb: &qb,
		mg: m,
	}
	return cf
}

func (f *ConnectionFactory) Connection() *Connection {
	return &Connection{
		db: f.db.Db,
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
