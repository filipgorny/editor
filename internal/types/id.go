package types

import "github.com/google/uuid"

type ID string

func NewId() ID {
	return ID(uuid.NewString())
}
