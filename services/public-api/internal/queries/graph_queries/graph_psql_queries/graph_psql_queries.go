package graph_psql_queries

import (
	"github.com/Masterminds/squirrel"
	"gopkg.in/gorp.v2"
	"github.com/gin-gonic/gin"

	"sg/services/public-api/internal/models"
)

type GraphPsqlQueries struct {
	db *gorp.DbMap
	qb *squirrel.StatementBuilderType
}

func (c *GraphPsqlQueries) Close() {
	c.db.Db.Close()
}

func (c *GraphPsqlQueries) DirectDB() *gorp.DbMap {
	return c.db
}

func (c *GraphPsqlQueries) CreateNode(ctx *gin.Context, v *models.Node) error {
	err := c.db.WithContext(ctx).Insert(v)
	if err != nil {
		return err
	}
	return nil
}

func (c *GraphPsqlQueries) GetNode(ctx *gin.Context, id string) (*models.Node, error) {
	query, args, err := c.qb.Select("*").From("nodes").Where("id=?", id).ToSql()
	if err != nil {
		return nil, err
	}

	v := &models.Node{}
	err = c.db.WithContext(ctx).SelectOne(v, query, args...)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (c *GraphPsqlQueries) CreateEdge(ctx *gin.Context, v *models.Edge) error {
	err := c.db.WithContext(ctx).Insert(v)
	if err != nil {
		return err
	}
	return nil
}

func (c *GraphPsqlQueries) GetEdge(ctx *gin.Context, id string) (*models.Edge, error) {
	query, args, err := c.qb.Select("*").From("edges").Where("id=?", id).ToSql()
	if err != nil {
		return nil, err
	}

	v := &models.Edge{}
	err = c.db.WithContext(ctx).SelectOne(v, query, args...)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (c *GraphPsqlQueries) CreateScope(ctx *gin.Context, v *models.Scope) error {
	err := c.db.WithContext(ctx).Insert(v)
	if err != nil {
		return err
	}
	return nil
}

func (c *GraphPsqlQueries) GetScope(ctx *gin.Context, id string) (*models.Scope, error) {
	query, args, err := c.qb.Select("*").From("scopes").Where("id=?", id).ToSql()
	if err != nil {
		return nil, err
	}

	v := &models.Scope{}
	err = c.db.WithContext(ctx).SelectOne(v, query, args...)
	if err != nil {
		return nil, err
	}

	return v, nil
}


