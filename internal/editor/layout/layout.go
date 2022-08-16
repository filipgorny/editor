package layout

import "github.com/filipgorny/editor/internal/editor/window"

type Layout struct {
	areas []*Area
}

func NewLayout() *Layout {
	return &Layout{
		areas: make([]*Area, 0),
	}
}

func (l *Layout) AddArea(area *Area) {
	l.areas = append(l.areas, area)
}

func (l *Layout) AttachWindow(name string, window *window.Window) {
	for _, area := range l.areas {
		if area.Name() == name {
			area.
		}
	}
}
