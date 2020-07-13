package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	port     = "5432"
	dockerUser = "postgres"
	password = "secret"
	name     = "postgres"
)

// TODO: currently is a global variable. Connection to db should be instantiated on a per-request basis.
var DB *sql.DB

func InitDb(host string) {
	user := dockerUser
	if host == "localhost" {
		// For postgres running in the background on local machine
		user = "stevenelleman"
	}

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
	fmt.Println("Successfully connected!")
}
