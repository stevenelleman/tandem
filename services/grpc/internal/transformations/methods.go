package transformations

// What's different between highlight and context? It's basically identical,
func Highlight(doc0 string, start, end int) string {
	return doc0[start:end]
}

// Handles additions, deletions, substitutions
func Replace(doc0, s string, start, end int) (string, string, int, int) {
	removed := doc0[start:end]
	inverseStart := start
	inverseEnd := start + len(s)
	return doc0[0:start] + s + doc0[end:], removed, inverseStart, inverseEnd
}

// Selected piece moved from one place to another
func Drag(doc0 string, start0, end0, start1 int) (string, int, int, int) {
	drag := doc0[start0:end0]
	start0Inverse := start1
	end0Inverse := start1 + len(drag)
	start1Inverse := start0
	doc1 := doc0[0:start0] + doc0[end0:]
	return doc1[0:start1] + drag + doc1[start1:], start0Inverse, end0Inverse, start1Inverse
}

// "Something given" vs "Something taken" -- Pull out a piece of a doc
func Take(doc0 string, start, end int) string {
	return doc0[start:end]
}

// Place source into doc
func Embed(doc0 string, s string, start, end int) (string, string, int, int) {
	inverseStart := start
	inverseEnd := start + len(s)
	return doc0[0:start] + s + doc0[end:], "", inverseStart, inverseEnd
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
// Or is this always available? Older parts of the graph are _always_ available?
// One consequence is that privacy is immutable -- once shared there's no going back
// In theory I like that -- there's no way to put the genie out of the bottle,
// might as well be explicit about it
func Import() {}

// TODO: How to create composeable functions gracefully? Difficult to anticipate necessary arguments.
// Logical Transformation Types:
func Transclude() {} // Embed(doc1, Take(doc0, "", args0), args1)
