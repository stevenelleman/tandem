package connection

import (
	"github.com/Masterminds/squirrel"
	"gopkg.in/gorp.v2"
)

type Connection struct {
	db *gorp.DbMap
	qb *squirrel.StatementBuilderType
}

func (c *Connection) Close() {
	c.db.Db.Close()
}
