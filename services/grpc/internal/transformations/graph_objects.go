package transformations

type Node struct {
	NodeId string `db:"nodeId"`
	CheckpointDoc string `db:"doc"` // Current state of the whole document, cached for some fraction of the nodes 
	docIsPopulated bool `db:"isDocPopulated"`

	EdgesOut string[] `db:"edgesOut"` // IDs of edges whose destination is this node
	EdgesIn string[] `db:"edgesIn"`   // ID of edges whose source is this node

	// Metadata
	Date string `db:"date"`
	Author string `db:"author"`
}

type Edge struct {
	EdgeId string `db:"edgeId"`
	SourceDoc string `db:"sourceDoc"`
	TargetDoc string `db:"targetDoc`
	Args FxnArgs `db:"args"` 				// f(SourceDoc, Args) -> TargetDoc
	InverseArgs FxnArgs `db:"inverseArgs"`  // f(TargetDoc, InverseArgs) -> SourceDoc

	SourceNodeId string `db:"sourceNodeId"`
	TargetNodeId string `db:"targetNodeId"`

	// Metadata
	Date string `db:"date"`
	Author string `db:"author"`
}

type Scope struct {
	ScopeId string `db:"scopeId"`
	
	// Scope consists of a set of nodes, and the set of edges between them
    Nodes string[] `db:"nodes"` // Node IDs
    Edges string[] `db:"edges"` // Edge IDs

	// Metadata struct, todo
	ScopeMetadata string[] `db:"scopeMetadata"`
}
