package editor

import (
	"github.com/filipgorny/editor/internal/editor/screen"
	"github.com/filipgorny/editor/internal/editor/workspace"
	"github.com/filipgorny/editor/internal/render"
)

type Editor struct {
	renderer render.Renderer
	screen   *screen.Screen
}

func NewEditor(renderer render.Renderer, workspace *workspace.Workspace) *Editor {
	return &Editor{
		renderer: renderer,
		screen:   screen.NewScreen(*renderer.Size(), workspace),
	}
}

func (e *Editor) Render() {
	for _, window := range e.screen.Workspace().Windows() {
		e.renderer.Display(*e.renderer.MaxRect(), window.DisplayContent())
	}
}
