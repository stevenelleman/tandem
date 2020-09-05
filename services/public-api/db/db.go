package db

import (
	"database/sql"
	"fmt"
	"libraries/datastore/migrater"

	_ "github.com/lib/pq"
)

const (
	port     = "5432"
	user     = "postgres"
	password = "secret"
	name     = "postgres"
)

// TODO: currently is a global variable. Connection to db should be instantiated on a per-request basis.
var DB *sql.DB

func InitDb(host string) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port,
		user, password, name)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Public-API Datastore: Successfully connected!")

	// Apply database migrations on Public-API Datastore
	m := migrater.Migrater{}
	count, err := m.Up(DB)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Public-API Datastore: Executed %d migrations\n", count)
}
