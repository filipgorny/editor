package events

import "github.com/filipgorny/editor/internal/types"

type Event interface {
	Id() *types.ID
}
