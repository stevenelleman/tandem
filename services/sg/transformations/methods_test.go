package transformations

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// go test fxns_test.go fxns.go -run TestReplace
// go test
func TestMethods(t *testing.T) {
	src1 := "hello..."

	// Addition
	doc1, isrc, istart, iend := Replace("", src1, 0, 0)
	require.Equal(t, src1, doc1)

	idoc1, src, start, end := Replace(doc1, isrc, istart, iend)
	require.Equal(t, "", idoc1)
	require.Equal(t, src1, src)
	require.Equal(t, 0, start)
	require.Equal(t, 0, end)

	// Deletion
	doc2, isrc, istart, iend := Replace(doc1, "", 0, 5)
	require.Equal(t, "...", doc2)

	idoc2, src, start, end := Replace(doc2, isrc, istart, iend)
	require.Equal(t, doc1, idoc2) // HERE: Failing
	require.Equal(t, "", src)
	require.Equal(t, 0, start)
	require.Equal(t, 5, end)

	// Deletion
	doc2, isrc, istart, iend = Replace(doc1, "", 5, 8)
	require.Equal(t, "hello", doc2)

	idoc2, src, start, end = Replace(doc2, isrc, istart, iend)
	require.Equal(t, doc1, idoc2)
	require.Equal(t, "", src)
	require.Equal(t, 5, start)
	require.Equal(t, 8, end)

	// Replace
	doc2, isrc, istart, iend = Replace(doc1, " hello", 5, 8)
	require.Equal(t, "hello hello", doc2)

	idoc2, src, start, end = Replace(doc2, isrc, istart, iend) // here
	require.Equal(t, doc1, idoc2)
	require.Equal(t, " hello", src)
	require.Equal(t, 5, start)
	require.Equal(t, 8, end)

	// Transclude -- Is not inverse but still returns initial state
	doc3, isrc, istart, iend := Transclude(doc2, " bye", 5, 5)
	require.Equal(t, "hello bye hello", doc3)

	idoc3, src, start, end := Transclude(doc3, isrc, istart, iend)
	require.Equal(t, doc2, idoc3)

	doc4, isrc, istart, iend := Transclude(doc3, " blah", 9, 9)
	require.Equal(t, "hello bye blah hello", doc4)

	idoc4, _, _, _ := Replace(doc4, isrc, istart, iend)
	require.Equal(t, doc3, idoc4)

	// Drag -- Seems to be inverse -- need to figure out logic of what is invertible
	doc5, istart0, iend0, istart1 := Drag(doc4, 6, 10, 11)
	require.Equal(t, "hello blah bye hello", doc5)

	idoc4, start0, end0, start1 := Drag(doc5, istart0, iend0, istart1)
	require.Equal(t, doc4, idoc4)
	require.Equal(t, 6, start0)
	require.Equal(t, 10, end0)
	require.Equal(t, 11, start1)
}
