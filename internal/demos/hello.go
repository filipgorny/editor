package demos

import (
	"log"

	"github.com/filipgorny/editor/internal/editor/buffer"
	"github.com/filipgorny/editor/internal/editor/editor"
	"github.com/filipgorny/editor/internal/editor/window"
	"github.com/filipgorny/editor/internal/editor/workspace"

	"github.com/filipgorny/editor/internal/driver/tcell_driver"
)

func Hello() {
	workspace := workspace.NewWorkspace()

	renderer := tcell_driver.NewTcellRender()

	buffer := buffer.NewBuffer()
	buffer.Content().WriteLine("Hello World!")
	buffer.Content().WriteLine("This is an example buffer")

	windowMain := window.NewWindow(buffer)

	workspace.AddWindow(windowMain)

	editor := editor.NewEditor(renderer, renderer, workspace)

	render := true

	if render {
		editor.Render()
	}

	editor.Run()

	log.Print()
}
