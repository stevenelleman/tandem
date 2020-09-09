package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	port     = "5432"
	user     = "postgres-user"
	password = "secret"
	name     = "postgresdb"
)

func InitDb(host string, conns int) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port,
		user, password, name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// TODO: tune
	db.SetMaxOpenConns(conns)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Public-API Datastore: Successfully connected!")

	return db
}
