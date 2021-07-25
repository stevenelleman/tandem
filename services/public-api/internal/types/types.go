package types

// All string document transformations
type StringType interface {
	Version() int
}

type StringDocument struct {
	Document string
}

// All type document transformations
type TypeType interface {
	StringType
}

type TypeDocument struct {}

// All scope document transformations
type ScopeType interface {
	StringType
}

type ScopeDocument struct {}

// All transformation document transformations
type TransformationType interface {
	StringType
}

type TransformationDocument struct {}