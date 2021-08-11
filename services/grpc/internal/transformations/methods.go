package transformations

// What's different between highlight and context? It's basically identical,
func Highlight(doc0 string, start, end int) string {
	return doc0[start:end]
}

// Handles additions, deletions, substitutions
func Replace(d, s string, start, end int) (string, string, int, int) {
	removed := d[start:end]
	inverseStart := start
	inverseEnd := start + len(s)
	return d[0:start] + s + d[end:], removed, inverseStart, inverseEnd
}

// Return new doc with embedded source
func Transclude(d, s string, start, end int) (string, string, int, int) {
	inverseStart := start
	inverseEnd := start + len(s)
	return d[0:start] + s + d[end:], "", inverseStart, inverseEnd
}

// Selected piece moved from one place to another
func Drag(d string, start0, end0, start1 int) (string, int, int, int) {
	drag := d[start0:end0]
	start0Inverse := start1
	end0Inverse := start1 + len(drag)
	start1Inverse := start0
	d1 := d[0:start0] + d[end0:]
	return d1[0:start1] + drag + d1[start1:], start0Inverse, end0Inverse, start1Inverse
}

// Other Transformation Types

// Mark the surrounding context for some highlight -- highlight of a highlight?
// Why can't this be done with highlight?
// Highlight the space around it, context link between the two
func Context() {}

// Unidirectional Link
func UniLink() {}

// Bidirectional Link -- use for branching
func BiLink() {}

// Intentionally link one region with another, any difference with Link? Denotes a logical linkage rather than a continuation of the document?
func Reference() {}

// Object is available in new context -- extremely important transformation
// Could import a document directly -- for a ML project for instance
// Import would make sense in a code context, not in a writing context
func Import() {}
