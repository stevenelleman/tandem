package graph_controller

import (
	"github.com/gin-gonic/gin"
	"sg/services/public-api/internal/models"
	"sg/services/public-api/internal/queries/graph_queries/graph_psql_queries"
)

type GraphController struct {
	query *graph_psql_queries.GraphPsqlQueries
}

func (c *GraphController) Close() {
	c.query.Close()
}

func (c *GraphController) GetNode(ctx *gin.Context, id string) (*models.Node, error) {
	return c.query.GetNode(ctx, id)
}

func (c *GraphController) CreateNode(ctx *gin.Context, s *models.Node) error {
	return c.query.CreateNode(ctx, s)
}

func (c *GraphController) GetEdge(ctx *gin.Context, id string) (*models.Edge, error) {
	return c.query.GetEdge(ctx, id)
}

func (c *GraphController) CreateEdge(ctx *gin.Context, s *models.Edge) error {
	return c.query.CreateEdge(ctx, s)
}

func (c *GraphController) GetScope(ctx *gin.Context, id string) (*models.Scope, error) {
	return c.query.GetScope(ctx, id)
}

func (c *GraphController) CreateScope(ctx *gin.Context, s *models.Scope) error {
	return c.query.CreateScope(ctx, s)
}


