package transformations

type Node struct {
	Id string `db:"id"`
	Doc string `db:"doc"`
}

type Edge struct {
	Id string `db:"id"`
	TxnId string `db:"txnId"`
	SourceDoc string `db:"sourceDoc"`
	InverseDoc string `db:"inverseDoc`
	Args FxnArgs `db:"args"`
	InverseArgs FxnArgs `db:"inverseArgs"`

	// Metadata
	Date string `db:"date"`
	Author string `db:"author"`
}

type Scope struct {
	Id string `db:"id"`
	
	// Scope consists of a set of nodes, and the set of edges between them
    Nodes string[] `db:"nodes"` // Node IDs
    Edges string[] `db:"edges"` // Edge IDs
}
