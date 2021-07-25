package transformations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// go test transformations_test.go transformations.go methods.go
func TestTransformations(t *testing.T) {
	src0 := "hello..."

	// Addition
	args := &FxnArgs{
		Source: src0,
	}
	doc1, iargs1 := Call("", args)
	require.Equal(t, src0, doc1)

	idoc1, iargs0 := Call(doc1, iargs1)
	require.Equal(t, "", idoc1)
	require.Equal(t, src0, iargs0.Source)
	require.Equal(t, 0, iargs0.OffsetStart1)
	require.Equal(t, 0, iargs0.OffsetEnd1)

	// Deletion
	args.Source = ""
	args.OffsetEnd1 = 5
	doc2, iargs2 := Call(doc1, args)
	require.Equal(t, "...", doc2)

	idoc2, iargs1 := Call(doc2, iargs2)
	require.Equal(t, doc1, idoc2)
	require.Equal(t, "", iargs1.Source)
	require.Equal(t, 0, iargs1.OffsetStart1)
	require.Equal(t, 5, iargs1.OffsetEnd1)

	// Deletion
	args.OffsetStart1 = 5
	args.OffsetEnd1 = 8
	doc2, iargs2 = Call(doc1, args)
	require.Equal(t, "hello", doc2)

	idoc2, iargs1 = Call(doc2, iargs2)
	require.Equal(t, doc1, idoc2)
	require.Equal(t, "", iargs1.Source)
	require.Equal(t, 5, args.OffsetStart1)
	require.Equal(t, 8, args.OffsetEnd1)

	// Replace
	args.Source = " hello"
	doc2, iargs2 = Call(doc1, args)
	require.Equal(t, "hello hello", doc2)

	idoc2, iargs1 = Call(doc2, iargs2)
	require.Equal(t, doc1, idoc2)
	require.Equal(t, " hello", iargs1.Source)
	require.Equal(t, 5, iargs1.OffsetStart1)
	require.Equal(t, 8, iargs1.OffsetEnd1)

	// Transclude -- Is not inverse but still returns initial state
	args.FxnId = EMBED
	args.OffsetEnd1 = 5
	args.Source = " bye"
	doc3, iargs2 := Call(doc2, args)
	require.Equal(t, "hello bye hello", doc3)

	idoc3, iargs1 := Call(doc3, iargs2)
	require.Equal(t, doc2, idoc3)

	args.OffsetStart1 = 9
	args.OffsetEnd1 = 9
	args.Source = " blah"
	doc4, iargs2 := Call(doc3, args)
	require.Equal(t, "hello bye blah hello", doc4)

	iargs2.FxnId = REPLACE
	idoc4, _ := Call(doc4, iargs2)
	require.Equal(t, doc3, idoc4)

	// Drag
	args.FxnId = DRAG
	args.OffsetStart1 = 6
	args.OffsetEnd1 = 10
	args.OffsetStart2 = 11
	args.Source = ""
	doc5, iargs2 := Call(doc4, args)
	require.Equal(t, "hello blah bye hello", doc5)

	idoc4, iargs1 = Call(doc5, iargs2)
	require.Equal(t, doc4, idoc4)
	require.Equal(t, 6, iargs1.OffsetStart1)
	require.Equal(t, 10, iargs1.OffsetEnd1)
	require.Equal(t, 11, iargs1.OffsetStart2)
}
