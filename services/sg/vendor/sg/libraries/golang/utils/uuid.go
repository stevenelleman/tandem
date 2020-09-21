package utils

import (
	"github.com/google/uuid"
)

func UUIDv4() uuid.UUID {
	u, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return u
}
