package models

// TODO: Remove when there are enough layering examples to remove silo examples
type Silo struct {
	Id string `db:"id"`
	State string `db:"state"` // TODO: Decide possible state values and replace with an enum. (active, disabled, deleted?)
}
