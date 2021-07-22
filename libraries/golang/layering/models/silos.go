package models

type Silo struct {
	Id string `db:"id"`
	State string `db:"state"` // TODO: Decide possible state values and replace with an enum. (active, disabled, deleted?)
}

type Node struct {
	Id string `db:"id"`
	Doc string `db:"doc"`
}

type Edge struct {
	Id string `db:"id"`
	Txn_id string `db:"txn_id"`
	Source_doc string `db:"source_doc"`
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