package editor

import (
	"os"

	"github.com/filipgorny/editor/internal/driver"
	"github.com/filipgorny/editor/internal/editor/events"
	"github.com/filipgorny/editor/internal/editor/screen"
	"github.com/filipgorny/editor/internal/editor/workspace"
	"github.com/filipgorny/editor/internal/log"
)

type Editor struct {
	renderer driver.Renderer
	sensor   driver.Sensor
	screen   *screen.Screen
}

func NewEditor(renderer driver.Renderer, sensor driver.Sensor, workspace *workspace.Workspace) *Editor {
	log.Log("creating editor")
	return &Editor{
		renderer: renderer,
		screen:   screen.NewScreen(*renderer.Size(), workspace),
		sensor:   sensor,
	}
}

func (e *Editor) Render() {
	log.Log("rendering editor")
	for _, window := range e.screen.Workspace().Windows() {
		log.Log("rendering window")
		e.renderer.Display(*e.renderer.MaxRect(), window.DisplayContent())
	}

	e.renderer.Sync()
}

func (e *Editor) Stop() {
	e.renderer.Stop()
	os.Exit(0)
}

func (editor *Editor) Run() {
	editor.sensor.Run(func(e events.Event) {
		switch e.Type() {
		case "key":
			key := e.(*events.KeyEvent).Key()

			if key == "q" {
				editor.Stop()
			}
		}
	})
}
