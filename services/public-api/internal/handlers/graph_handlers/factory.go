package graph_handlers

import (
	"sg/services/public-api/internal/controllers/graph_controller"
	"sg/services/public-api/internal/queries/graph_queries/graph_psql_queries"
)

func NewHandler(psqlArgs *graph_psql_queries.StoreArgs) *GraphHandlers {
	cf := graph_controller.NewControllerFactory(psqlArgs)
	return &GraphHandlers{
		controller: cf.Controller(),
	}
}