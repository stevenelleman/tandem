package db

import (
	"database/sql"
	"fmt"
	"sg/services/public-api/constants"

	_ "github.com/lib/pq"
)

func InitDbConn(host string, conns int) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host,
		constants.APIStorePort,
		constants.APIStoreUser,
		constants.APIStorePassword,
		constants.APIStoreName,
	)

	db, err := sql.Open(constants.APIStoreDriver, psqlInfo)
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
