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
	Args string `db:"args"`

	// Metadata
	Date string `db:"date"`
	Author string `db:"author"`
}

type Scope struct {
	Id string `db:"id"`
	
	// Scope consists of a set of node IDs, and the set of Edge IDs between those nodes
    Nodes string[] `db:"nodes"`
    Edges string[] `db:"edges"`
}
