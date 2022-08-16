package workspace

import "github.com/filipgorny/editor/internal/editor/window"

type Workspace struct {
	windowsList *window.WindowsList
}

func NewWorkspace() *Workspace {
	return &Workspace{
		windowsList: window.NewWindowsList(),
	}
}

func (w *Workspace) Windows() []*window.Window {
	return w.windowsList.Windows()
}

func (w *Workspace) AddWindow(window *window.Window) {
	w.windowsList.AddWindow(window)
}
