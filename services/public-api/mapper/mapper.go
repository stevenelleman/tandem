package mapper

import (
	"web-microservice-shell/libraries/golang/guts/models"

	"gopkg.in/gorp.v2"
)

// All table primary keys should be id
func MapTables(dbmap *gorp.DbMap) {
	silos := dbmap.AddTableWithName(models.Silo{}, "silos")
	silos.SetKeys(false, "id")
}
