package window

type WindowsList struct {
	windows []*Window
	active  *Window
}

func NewWindowsList() *WindowsList {
	return &WindowsList{
		windows: make([]*Window, 0),
		active:  nil,
	}
}

func (w *WindowsList) AddWindow(window *Window) {
	w.windows = append(w.windows, window)

	if w.active == nil {
		w.active = window
	}
}

func (w *WindowsList) Active() *Window {
	return w.active
}

func (w *WindowsList) SetActive(window *Window) {
	w.active = window
}

func (w *WindowsList) Windows() []*Window {
	return w.windows
}
