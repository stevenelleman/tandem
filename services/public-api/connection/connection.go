package connection

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
)

type Connection struct {
	db *sql.DB
	qb *squirrel.StatementBuilderType
}

func (c *Connection) Close() {
	c.db.Close()
}
