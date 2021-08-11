package transformations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// go test transformations_test.go transformations.go methods.go
func TestTransformations(t *testing.T) {
	src0 := "hello..."

	// Addition
	args := &FxnArgs{}
	doc1, isrc1, iargs1 := Call("", src0, args)
	require.Equal(t, src0, doc1)

	idoc1, isrc0, iargs0 := Call(doc1, isrc1, iargs1)
	require.Equal(t, "", idoc1)
	require.Equal(t, src0, isrc0)
	require.Equal(t, 0, iargs0.OffsetStart1)
	require.Equal(t, 0, iargs0.OffsetEnd1)

	// Deletion
	args.OffsetEnd1 = 5
	doc2, isrc2, iargs2 := Call(doc1, "", args)
	require.Equal(t, "...", doc2)

	idoc2, isrc1, iargs1 := Call(doc2, isrc2, iargs2)
	require.Equal(t, doc1, idoc2)
	require.Equal(t, "", isrc1)
	require.Equal(t, 0, iargs1.OffsetStart1)
	require.Equal(t, 5, iargs1.OffsetEnd1)

	// Deletion
	args.OffsetStart1 = 5
	args.OffsetEnd1 = 8
	doc2, isrc2, iargs2 = Call(doc1, "", args)
	require.Equal(t, "hello", doc2)

	idoc2, isrc1, iargs1 = Call(doc2, isrc2, iargs2)
	require.Equal(t, doc1, idoc2)
	require.Equal(t, "", isrc1)
	require.Equal(t, 5, args.OffsetStart1)
	require.Equal(t, 8, args.OffsetEnd1)

	// Replace
	doc2, isrc2, iargs2 = Call(doc1, " hello", args)
	require.Equal(t, "hello hello", doc2)

	idoc2, isrc1, iargs1 = Call(doc2, isrc2, iargs2)
	require.Equal(t, doc1, idoc2)
	require.Equal(t, " hello", isrc1)
	require.Equal(t, 5, iargs1.OffsetStart1)
	require.Equal(t, 8, iargs1.OffsetEnd1)

	// Transclude -- Is not inverse but still returns initial state
	args.FxnId = TRANSCLUDE
	args.OffsetEnd1 = 5
	doc3, isrc2, iargs2 := Call(doc2, " bye", args)
	require.Equal(t, "hello bye hello", doc3)

	idoc3, isrc1, iargs1 := Call(doc3, isrc2, iargs2)
	require.Equal(t, doc2, idoc3)

	args.OffsetStart1 = 9
	args.OffsetEnd1 = 9
	doc4, isrc2, iargs2 := Call(doc3, " blah", args)
	require.Equal(t, "hello bye blah hello", doc4)

	iargs2.FxnId = REPLACE
	idoc4, _, _ := Call(doc4, isrc2, iargs2)
	require.Equal(t, doc3, idoc4)

	// Drag
	args.FxnId = DRAG
	args.OffsetStart1 = 6
	args.OffsetEnd1 = 10
	args.OffsetStart2 = 11
	doc5, isrc2, iargs2 := Call(doc4, "", args)
	require.Equal(t, "hello blah bye hello", doc5)

	idoc4, _, iargs1 = Call(doc5, isrc2, iargs2)
	require.Equal(t, doc4, idoc4)
	require.Equal(t, 6, iargs1.OffsetStart1)
	require.Equal(t, 10, iargs1.OffsetEnd1)
	require.Equal(t, 11, iargs1.OffsetStart2)
}
