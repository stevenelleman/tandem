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

// Return new doc with embedded source
// A: Start is inclusive
func Transclude(d, s string, start int) string {
	return d[0:start] + s + d[start:]
}

// Return new source that's attached to the doc, the highlight is _part_ of the doc
// A: Start and end are inclusive
func Highlight(doc0 string, start, end int) string {
	return doc0[start:end]
}

// Selected piece moved from one place to another
// start1 calculated after deletion
func Drag(d, s string, start0, start1 int) string {
	d1 := d[0:start0] + d[start0+len(s):]
	return d1[0:start1] + s + d1[start1:]
}

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
