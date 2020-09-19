package main

import (
	"sg/libraries/golang/guts/models"

	"gopkg.in/gorp.v1"
)

// Map golang model to postgres table
// All table primary keys should be id
// TODO: better code reuse?
func mapTables(dbmap *gorp.DbMap) {
	silos := dbmap.AddTableWithName(models.Silo{}, "silos")
	silos.SetKeys(false, "id")
}
