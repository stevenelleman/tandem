package fxns

/*
TODO: One directory for function, each function gets a directory
	Use the same format, similar to a sql migration file
	Rigorous tests in each directory
	Create script outside directories that:
	1. Gets directory names
	2. Generates constants based on those names
	3. All functions should take: region0, src0, args and return region1, iregion, iargs
	4. Just regenerating the Call method would likely be sufficient to continue.
		At that point, the fxn id will be available for use, as well as special arguments in FxnArgs
	5. Would be even cooler if from the frontend UI could create this directory.
		Have an internal review process? Too expensive.

Notes:
- would be cool if functions were only code local to sg, that controller only used fxns.Call
	and then everything else was completely generalized for sources, metadata, scopes and commits
	that would be pretty cool -- it would really minimize changes to the controller code, which would be terrific
	could go meta and have a transformation client, which is basically the one call -- not a bad idea
	to use the same controller-client pattern
*/

// Handles additions, deletions, substitutions
func Replace(d, s string, start, end int) string {
	return d[0:start] + s + d[end:]
}

// Return new doc with embedded source
func Transclude(d, s string, start int) string {
	return d[0:start] + s + d[start:]
}

// Selected piece moved from one place to another
// start1 calculated after deletion
func Drag(d, s string, start0, start1 int) string {
	d1 := d[0:start0] + d[start0+len(s):]
	return d1[0:start1] + s + d1[start1:]
}

// Return new source that's attached to the doc, the highlight is _part_ of the doc
// Inverse represented by "unmaking" source, which "unasks" the question
// TODO: Need to look at highlighter code, merge code is important
func Highlight(doc0 string, start, end int) string {
	return doc0[start:end]
}

// Inverse methods produce arguments to move back to the original source
// Inverse arguments make not produce original arguments -- i.e. function may not be completely invertible

// Affected region should be start to (start + len(s))
func ReplaceInverse(d, s string, start, end int) (string, string, int, int) {
	removed := d[start:end]
	inverseStart := start
	inverseEnd := start + len(s)
	return d[0:start] + s + d[end:], removed, inverseStart, inverseEnd
}

// Affected region should be start to (start + len(s))
func TranscludeInverse(d, s string, start, end int) (string, string, int, int) {
	inverseStart := start
	inverseEnd := start + len(s)
	return d[0:start] + s + d[end:], "", inverseStart, inverseEnd
}

// Affected region should be (start0, end0) and (start1, start1 + (end0 - start0))
func DragInverse(d string, start0, end0, start1 int) (string, int, int, int) {
	drag := d[start0:end0]

	start0Inverse := start1
	end0Inverse := start1 + len(drag)
	start1Inverse := start0
	d1 := d[0:start0] + d[end0:]
	return d1[0:start1] + drag + d1[start1:], start0Inverse, end0Inverse, start1Inverse
}

// Other Transformation Types

// Mark the surrounding context for some highlight -- highlight of a highlight?
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
