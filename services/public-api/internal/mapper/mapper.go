package mapper

import (

	"gopkg.in/gorp.v2"
	"sg/services/public-api/internal/models"
	example_models "sg/libraries/golang/layering/models"
)

// All table primary keys should be id
func MapTables(dbmap *gorp.DbMap) {
	silos := dbmap.AddTableWithName(models.Silo{}, "silos")
	silos.SetKeys(false, "id")
}
