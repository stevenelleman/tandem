package fxns

/*
V0 Transformations: Basic Math [V1: Data Structures]
Objective: Offer "functional diffs" that can build a text document
- Need some kind of mapping between data and ui representation

Resources:
- https://yourbasic.org/golang/build-append-concatenate-strings-efficiently/
*/

// Handles additions, deletions, substitutions
func Replace(d, s string, start, end int) string {
	return d[0:start] + s + d[end:]
}

// Return new doc with embedded source
func Transclude(d, s string, start int) string {
	return d[0:start] + s + d[start:]
}

// Return new source that's attached to the doc, the highlight is _part_ of the doc
// Inverse represented by "unmaking" source, which "unasks" the question
func Highlight(doc0 string, start, end int) string {
	return doc0[start:end]
}

// Selected piece moved from one place to another
// start1 calculated after deletion
func Drag(d, s string, start0, start1 int) string {
	d1 := d[0:start0] + d[start0+len(s):]
	return d1[0:start1] + s + d1[start1:]
}

// Inverse methods produce arguments to move back to the original source
// Inverse arguments make not produce original arguments -- i.e. function may not be completely invertible

func ReplaceInverse(d, s string, start, end int) (string, string, int, int) {
	removed := d[start:end]
	inverseStart := start
	inverseEnd := start + len(s)
	return d[0:start] + s + d[end:], removed, inverseStart, inverseEnd
}

func TranscludeInverse(d, s string, start, end int) (string, string, int, int) {
	inverseStart := start
	inverseEnd := start + len(s)
	return d[0:start] + s + d[end:], "", inverseStart, inverseEnd
}

func DragInverse(d string, start0, end0, start1 int) (string, int, int, int) {
	drag := d[start0:end0]

	start0Inverse := start1
	end0Inverse := start1 + len(drag)
	start1Inverse := start0
	d1 := d[0:start0] + d[end0:]
	return d1[0:start1] + drag + d1[start1:], start0Inverse, end0Inverse, start1Inverse
}

// Other Transformation Types

// Mark the surrounding context for some highlight
func Context() {}

// Start a new source, references back to the document
// What would this look like?
func Branch() {}

// Bidirectional link between one region and another, link has primary direction
func Link() {}

// Intentionally link one region with another, any difference with Link?
func Reference() {}

// Object is available in new context -- extremely important transformation
// Could import a document directly -- for a ML project for instance
// Import would make sense in a code context, not in a writing context
func Import() {}
