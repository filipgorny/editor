package layout

import (
	"github.com/filipgorny/editor/internal/editor/view"
	"github.com/filipgorny/editor/internal/types"
)

type Area struct {
	name    string
	rect    *types.Rect
	content *view.DisplayContent
}

func NewArea(name string, rect *types.Rect) *Area {
	return &Area{
		name:    name,
		rect:    rect,
		content: view.NewDisplayContent(rect.Width(), rect.Height()),
	}
}

func (a *Area) Name() string {
	return a.name
}

func (a *Area) Rect() *types.Rect {
	return a.rect
}
