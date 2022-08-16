package render

import (
	"github.com/filipgorny/editor/internal/editor/view"
	"github.com/filipgorny/editor/internal/types"
	"github.com/zyedidia/tcell/v2"
)

type TcellRender struct {
	screen tcell.Screen
	style  tcell.Style
}

func NewTcellRender() *TcellRender {
	screen, _ := tcell.NewScreen()

	return &TcellRender{
		screen: screen,
		style:  tcell.StyleDefault,
	}
}

func (r *TcellRender) Size() *types.Size {
	x, y := r.screen.Size()

	return types.NewSize(x, y)
}

func (r *TcellRender) SetContent(x, y int, ch rune) {
	r.screen.SetContent(x, y, ch, nil, r.style)
}

func (r *TcellRender) Display(rect types.Rect, content *view.DisplayContent) {
	for y := 0; y < rect.Height(); y++ {
		if y > rect.Height() {
			break
		}

		for x := 0; x < rect.Width(); x++ {
			if x > rect.Width() {
				break
			}

			r.SetContent(rect.X()+x, rect.Y()+y, content.Get(x, y))
		}
	}
}

func (r *TcellRender) MaxRect() *types.Rect {
	return types.NewRect(0, 0, r.Size().Width(), r.Size().Height())
}
