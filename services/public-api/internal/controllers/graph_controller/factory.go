package graph_controller

import "sg/services/public-api/internal/queries/graph_queries/graph_psql_queries"

type GraphControllerFactory struct {
	psqlConnFactory *graph_psql_queries.GraphPsqlConnectionFactory
}

func NewControllerFactory(psqlArgs *graph_psql_queries.StoreArgs) *GraphControllerFactory {
	cf := &GraphControllerFactory{}
	cf.psqlConnFactory = graph_psql_queries.NewPsqlConnFactory(psqlArgs)
	return cf
}

func (f *GraphControllerFactory) Controller() *GraphController {
	return &GraphController{
		query: f.psqlConnFactory.Connection(),
	}
}