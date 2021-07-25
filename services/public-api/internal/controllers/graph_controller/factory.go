package graph_controller

import "sg/services/public-api/internal/queries/graph_queries/psql"

type GraphControllerFactory struct {
	psqlConnFactory *psql.GraphPsqlConnectionFactory
}

func NewControllerFactory(psqlArgs *psql.StoreArgs) *GraphControllerFactory {
	cf := &GraphControllerFactory{}
	cf.psqlConnFactory = psql.NewPsqlConnFactory(psqlArgs)
	return cf
}

func (f *GraphControllerFactory) Controller() *GraphController {
	return &GraphController{
		query: f.psqlConnFactory.Connection(),
	}
}