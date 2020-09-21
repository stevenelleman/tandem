package models

import (
	"sg/libraries/golang/utils"
	"time"

	"github.com/google/uuid"
)

/*
// TODO: Neo4j queries -- are these correct atomic actions?
SG Connection Methods:
- GetSource
- ListSources
	- Need to grab sources and cache them
- CreateSource
- GetCommit
- ListCommits
	- Need to grab commits and cache them
- CreateCommit
- GetScope
- ListScopes
- CreateScope
- UpdateScope

Unresolved:
- UpdateSource // Should only be for emergencies, even then better to not have
- DeleteScope // Maybe need something to "unmake" Scope. better to simply make it private


*/

// Immutable
type Source struct {
	Id       uuid.UUID // ID of source, primary key
	AuthorId uuid.UUID // ID of author that created source
	Contents string

	//
	CreatedAt time.Time
}

func MakeSource(c string, authorId uuid.UUID) *Source {
	return &Source{
		Id:        utils.UUIDv4(),
		AuthorId:  authorId,
		Contents:  c,
		CreatedAt: time.Now().UTC(),
	}
}

// Small enough to save all versions
// Should be extensible
// Should be used for sorting purposes -- would require some kind of join between commits and metadata. Need to think about it
type Metadata struct {
	Id     uuid.UUID
	Labels []string // In list of key-value pairs or format: "key:value"
	Tags   []string //
}

// Immutable -- no updated_at or deleted_at
// Emphasize on lineages/phylogeny would be nice
// Should the commit express the explicit range it's affecting?
type Commit struct {
	Index    int       // Commit number
	Id       uuid.UUID // ID of version, primary key
	AuthorId uuid.UUID // ID of author that created scope node

	// Spine
	ScopeId  uuid.UUID // ID of scope
	ParentId uuid.UUID // ID of last Commit
	ChildId  uuid.UUID // ID of next Commit

	// Contents
	MetadataId  bool      // Metadata of Scope
	SourceId    uuid.UUID // ID of applied source
	Args        string    // Args -- not saved as structured data -- destructure it
	InverseArgs string    // Calculated inverse args

	// Side Info
	CreatedAt time.Time
}

// Where does metadata go? Would like to subject metadata changes to the same
// 	process as other elements of doc, but you likely need metadata
// Mutable
type Scope struct {
	Id       uuid.UUID
	AuthorId uuid.UUID

	FirstCommitId uuid.UUID
	LastCommitId  uuid.UUID

	LastCheckpointId uuid.UUID
}
