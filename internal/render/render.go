package render

import (
	"github.com/filipgorny/editor/internal/editor/view"
	"github.com/filipgorny/editor/internal/types"
)

type Renderer interface {
	Size() *types.Size
	SetContent(x, y int, ch rune)
	Display(rect types.Rect, content *view.DisplayContent)
	MaxRect() *types.Rect
}
