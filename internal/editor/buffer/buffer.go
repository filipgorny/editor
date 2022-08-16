package buffer

import (
	"github.com/filipgorny/editor/internal/types"
)

type Buffer struct {
	content *Content
	id      types.ID
}

func NewBuffer() *Buffer {
	return &Buffer{
		content: NewContent(),
		id:      types.NewId(),
	}
}

func (b *Buffer) Id() types.ID {
	return b.id
}

func (b *Buffer) Content() *Content {
	return b.content
}
