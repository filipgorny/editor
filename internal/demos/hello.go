package demos

import (
	"github.com/filipgorny/editor/internal/editor/buffer"
	"github.com/filipgorny/editor/internal/editor/editor"
	"github.com/filipgorny/editor/internal/editor/window"
	"github.com/filipgorny/editor/internal/editor/workspace"
	"github.com/filipgorny/editor/internal/render"
)

func Hello() {
	workspace := workspace.NewWorkspace()

	renderer := render.NewTcellRender()

	buffer := buffer.NewBuffer()
	buffer.Content().WriteLine("Hello World!")

	windowMain := window.NewWindow(buffer)

	workspace.AddWindow(windowMain)

	editor := editor.NewEditor(renderer, workspace)

	editor.Render()
}
