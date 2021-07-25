package models

import (
	"time"
	"web.microservice.shell/libraries/golang/utils"
)

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

// Create the fist node in a document, with no source edges
func ConstructNode(doc string, parent *Node) (*Node) {
	return &Node {
		Id: utils.UUIDv4().String(),
		Document: doc,
		Scopes: parent.Scopes, // Retain the same scopes as parent

		CreatedAt: time.Now(), // TODO: Express as a PreInsert so db clock is source of truth
	}
}

type FxnArgs struct {
	FxnId        int
	Source string
	OffsetStart1 int
	OffsetEnd1   int
	OffsetStart2 int
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

func ConstructEdge(args, iargs *FxnArgs, parentId, childId string) (*Edge) {
	return &Edge{
		Id: utils.UUIDv4().String(),

		Args: args,
		InverseArgs: iargs,

		ParentNodeId: parentId,
		ChildNodeId: childId,

		CreatedAt: time.Now(), // TODO: Express as a PreInsert so db clock is source of truth
	}
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
