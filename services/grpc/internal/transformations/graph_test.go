package transformations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// go test graph_objects_test.go graph_objects.go transformations.go methods.go
// TODO solidify the interface for Call()
func TestGraphObjects(t *testing.T) {
	doc0 := "graph test ..."

	// Addition
	args0 := &FxnArgs{}
	doc1, idoc1, iargs1 := Call(" and friends", doc0, args0)
	require.Equal(t, doc1, "graph test ... and friends")

	doc2, idoc2, iargs2 := Call(doc1, idoc1, iargs1)
	require.Equal(t, doc2, doc0)
	require.Equal(t, doc2, "graph test ...")

	// Create initial node
	var nodeId int = 0
	var node0 Node* = Createnode0(doc0, nodeId)
	require.Equal(t, node0.nodeId, 0)
	require.Equal(t, node0.CheckpointDoc, doc0)

	// Create an edge with a transformation
	var node1 Node*, edge01 *Edge = ApplyTransformation(node0, args0, doc0)
	// TODO: results
}
