package fxns

/*
V0 Transformations: Basic Math [V1: Data Structures]
Objective: Offer "functional diffs" that can build a text document
- Need some kind of mapping between data and ui representation

Resources:
- https://yourbasic.org/golang/build-append-concatenate-strings-efficiently/
*/

// Handles additions, deletions, substitutions
// A: Start and end are inclusive
func Replace(d, s string, start, end int) string {
	return d[0:start] + s + d[end:]
}

// TODO: Add affected region in both directions, that way the
// 	rest of the doc doesn't need to be transformed
// How to have multiple affected regions?
// Replace plus inverse arguments
func ReplaceInverse(d, s string, start, end int) (string, string, int, int) {
	removed := d[start:end]
	inverseStart := start
	inverseEnd := start + len(s)
	return d[0:start] + s + d[end:], removed, inverseStart, inverseEnd
}

// Return new doc with embedded source
// A: Start is inclusive
func Transclude(d, s string, start int) string {
	return d[0:start] + s + d[start:]
}

// Cannot capture inverse because there's no deletion
func TranscludeInverse(d, s string, start, end int) (string, string, int, int) {
	inverseStart := start
	inverseEnd := start + len(s)
	return d[0:start] + s + d[end:], "", inverseStart, inverseEnd
}

// Return new source that's attached to the doc, the highlight is _part_ of the doc
// A: Start and end are inclusive
// HERE: How to represent inverse? Matter of turning some kind of source characteristic back
// alternatively could make an "unmake" source, which "unasks" the question
func Highlight(doc0 string, start, end int) string {
	return doc0[start:end]
}

// Selected piece moved from one place to another
// start1 calculated after deletion
func Drag(d, s string, start0, start1 int) string {
	d1 := d[0:start0] + d[start0+len(s):]
	return d1[0:start1] + s + d1[start1:]
}

func DragInverse(d string, start0, end0, start1 int) (string, int, int, int) {
	drag := d[start0:end0]

	start0Inverse := start1
	end0Inverse := start1 + len(drag)
	start1Inverse := start0
	d1 := d[0:start0] + d[end0:]
	return d1[0:start1] + drag + d1[start1:], start0Inverse, end0Inverse, start1Inverse
}

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
