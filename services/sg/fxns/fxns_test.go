package fxns

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

//TODO: add benchmarking tests (https://stackoverflow.com/questions/24893624/how-to-replace-a-letter-at-a-specific-index-in-a-string-in-go)

// go test fxns_test.go fxns.go -run TestReplace
// go test
func TestBasics(t *testing.T) {
	src1 := "hello..."
	doc1 := Replace("", src1, 0, 0)
	require.Equal(t, src1, doc1)

	doc2 := Replace(doc1, "", 0, 5)
	require.Equal(t, "...", doc2)

	doc2 = Replace(doc1, "", 5, 8)
	require.Equal(t, "hello", doc2)

	doc2 = Replace(doc1, " hello", 5, 8)
	require.Equal(t, "hello hello", doc2)

	doc3 := Transclude(doc2, " bye", 5)
	require.Equal(t, "hello bye hello", doc3)

	doc3 = Transclude(doc3, " blah", 9)
	require.Equal(t, "hello bye blah hello", doc3)

	doc4 := Drag(doc3, "bye ", 6, 11)
	require.Equal(t, "hello blah bye hello", doc4)

	doc5 := Highlight(doc4, 6, 10)
	require.Equal(t, "blah", doc5)
}
