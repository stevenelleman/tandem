package models

import "github.com/google/uuid"

// Going to likely have to cache documents by session
// Can't be reconstructing docs from scratch all the time, need to start from some checkpoiunt
// Most recent version of a doc should have a checkpoint.
// Beginning needs no checkpoint, just start from zero
type Checkpoint struct {
	Index    int // Index of checkpoint
	SourceId uuid.UUID
	CommitId uuid.UUID
	Document string // Full document for this particular index
	// Where to handle chunking? -- probably in the query itself
}
