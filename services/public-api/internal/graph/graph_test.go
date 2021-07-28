package graph

import (
	"sg/services/public-api/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
)

// go test graph_objects_test.go graph_objects.go transformations.go methods.go
// TODO solidify the interface for Call()
func TestGraphObjects(t *testing.T) {
	doc0 := "graph test ..."
	added := " and friends"
	expected := "graph test ... and friends"

	// Addition
	args0 := &models.FxnArgs{
		FxnId: REPLACE,
		Source: added,
		OffsetStart1: len(doc0),
		OffsetEnd1: len(doc0),
	}
	doc1, iargs1 := Call(doc0, args0)
	require.Equal(t, doc1, expected)

	doc2, iargs2 := Call(doc1, iargs1)
	require.Equal(t, doc0, doc2)
	require.Equal(t, args0.Source, iargs2.Source)

	// Create initial node
	node0 := models.ConstructNode(doc0, &models.Node{})
	require.Equal(t, node0.Document, doc0)

	// Apply transformation
	node1, edge := Transform(node0, args0)
	require.Equal(t, expected, node1.Document)
	require.Equal(t, args0, edge.Args)
	require.Equal(t, iargs1, edge.InverseArgs)

	// Mock edits
	// 1. Each transform will need the current latest Node -- consider just needing the current document version
}
