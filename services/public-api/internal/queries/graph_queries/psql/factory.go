package psql

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"gopkg.in/gorp.v2"

	"sg/libraries/golang/datastore/migrater"
)

type StoreArgs struct {
	driver       string
	host         string
	user         string
	password     string
	dbname       string
	port         int
	maxConns     int
	migraterArgs *migrater.Args
	mapper       func(*gorp.DbMap)
}

func MakeArgs(driver, host, user, password, dbname string, port, conns int, migraterArgs *migrater.Args, mapper func(*gorp.DbMap)) *StoreArgs {
	return &StoreArgs{
		driver:       driver,
		host:         host,
		user:         user,
		password:     password,
		dbname:       dbname,
		port:         port,
		maxConns:     conns,
		migraterArgs: migraterArgs,
		mapper:       mapper,
	}
}

func InitDbConn(args *StoreArgs) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		args.host,
		args.port,
		args.user,
		args.password,
		args.dbname,
	)

	db, err := sql.Open(args.driver, psqlInfo)
	if err != nil {
		panic(err)
	}

	// TODO: tune
	db.SetMaxOpenConns(args.maxConns)

	err = db.Ping()
	fmt.Println("Ping", err)

	if err != nil {
		panic(err)
	}
	fmt.Println("Public-API Datastore: Successfully connected!")

	return db
}

type GraphPsqlConnectionFactory struct {
	db *gorp.DbMap
	qb *squirrel.StatementBuilderType
}

func NewPsqlConnFactory(args *StoreArgs) *GraphPsqlConnectionFactory {
	conn := InitDbConn(args)
	qb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// Apply database migrations on Public-API Datastore
	m := migrater.MakeMigrater(args.migraterArgs)
	count, err := m.Up(conn)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Public-API Datastore: Executed %d migrations\n", count)

	dbMap := &gorp.DbMap{Db: conn, Dialect: gorp.PostgresDialect{}}
	args.mapper(dbMap)

	cf := &GraphPsqlConnectionFactory{
		db: dbMap,
		qb: &qb,
	}
	return cf
}

// Note: in the future if a subset of methods can be reused make a separate interface for them and move them out into a separate library,
//	then compose them into the final connection type interface to reuse the logic. The same pattern goes for controller and handler levels.
//	In the future I can forsee this kind of reuseability applying to profile / auth logic
func (f *GraphPsqlConnectionFactory) Connection() *GraphPsqlQueries {
	return &GraphPsqlQueries{
		db: f.db,
		qb: f.qb,
	}
}
