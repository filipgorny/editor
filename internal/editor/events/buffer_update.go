package events

import (
	"github.com/filipgorny/editor/internal/editor/buffer"
	"github.com/filipgorny/editor/internal/types"
)

type BufferUpdate struct {
	Id       *types.ID
	BufferId *types.ID
}

func NewBufferUpdate(buffer *buffer.Buffer) *BufferUpdate {
	return &BufferUpdate{
		Id:       types.NewId(),
		BufferId: buffer.Id(),
	}
}
