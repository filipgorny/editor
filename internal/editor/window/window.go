package window

import (
	"github.com/filipgorny/editor/internal/editor/buffer"
	"github.com/filipgorny/editor/internal/editor/view"
	"github.com/filipgorny/editor/internal/types"
)

type Window struct {
	rect   *types.Rect
	id     types.ID
	buffer *buffer.Buffer
}

func NewWindow(buffer *buffer.Buffer) *Window {
	return &Window{
		rect:   types.NewRect(0, 0, 10, 10),
		id:     types.NewId(),
		buffer: buffer,
	}
}

func (w *Window) Id() types.ID {
	return w.id
}

func (w *Window) Rect() *types.Rect {
	return w.rect
}

func (w *Window) DisplayContent() *view.DisplayContent {
	dc := view.NewDisplayContent(w.Rect().Width(), w.Rect().Height())

	for i := 0; i < w.Rect().Height(); i++ {
		dc.Write(i, w.buffer.Content().Line(i))
	}

	return dc
}
