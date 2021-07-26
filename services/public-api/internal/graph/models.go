package graph

import (
	"time"
	"sg/libraries/golang/utils"
)

// Structs used in the graph implementation

type FxnArgs struct {
	FxnId        int
	Source       string
	OffsetStart1 int
	OffsetEnd1   int
	OffsetStart2 int
}

type Node struct {
	Id string `db:"id"`

	// TODO: Will Node -> Edge -> Node linked list be performant?
	EdgesOut []string `db:"edges_out"` // IDs of edges whose destination is this node
	EdgeIn string `db:"edge_in"`   // ID of edges whose source is this node
	Scopes []string `db:"scopes"` // IDs of scopes this node belongs to

	// TODO: Add chunk and type logic.
	Document string `db:"document"`

	// Metadata -- TODO: versioned metadata bytes
	CreatedAt time.Time `db:"created_at"`
	Author string `db:"author"`
}

type Edge struct {
	Id string `db:"id"`

	// TODO: versioned + typed FxnArgs
	Args *FxnArgs `db:"args"`                // f(SourceDoc, Args) -> TargetDoc
	InverseArgs *FxnArgs `db:"inverse_args"`  // f(TargetDoc, InverseArgs) -> SourceDoc

	ParentNodeId string `db:"parent_node_id"`
	ChildNodeId string `db:"child_node_id"`

	// Metadata
	CreatedAt time.Time `db:"created_at"`
	Author string `db:"author"`
}

// TODO: Think about how to represent Scope as a document itself
type Scope struct {
	Id string `db:"id"`

	// Scope consists of a set of nodes, and the set of edges between them
	// The set of edges is defined by the set of nodes
    Nodes []string `db:"node_ids"` // Node IDs

	// Metadata
	CreatedAt time.Time `db:"created_at"`
	Author string `db:"author"`
}
