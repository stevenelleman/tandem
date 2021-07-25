package graph_handlers

import (
	"sg/services/public-api/internal/controllers/graph_controller"
	"sg/services/public-api/internal/queries/graph_queries/psql"
)

func NewHandler(psqlArgs *psql.StoreArgs) *GraphHandlers {
	cf := graph_controller.NewControllerFactory(psqlArgs)
	return &GraphHandlers{
		controller: cf.Controller(),
	}
}