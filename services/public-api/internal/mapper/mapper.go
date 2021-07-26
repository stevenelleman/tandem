package mapper

import (

	"gopkg.in/gorp.v2"
	"sg/services/public-api/internal/models"
	example_models "sg/libraries/golang/layering/models"
)

// All table primary keys should be id
func MapTables(dbmap *gorp.DbMap) {
	silo := dbmap.AddTableWithName(example_models.Silo{}, "silos")
	silo.SetKeys(false, "id")

	nodes := dbmap.AddTableWithName(models.Node{}, "nodes")
	nodes.SetKeys(false, "id")

	edges := dbmap.AddTableWithName(models.Edge{}, "edges")
	edges.SetKeys(false, "id")

	scopes := dbmap.AddTableWithName(models.Scope{}, "scopes")
	scopes.SetKeys(false, "id")
}
