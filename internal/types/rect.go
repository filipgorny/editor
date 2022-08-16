package types

type Rect struct {
	x      int
	y      int
	width  int
	height int
}

func NewRect(x int, y int, width int, height int) *Rect {
	return &Rect{
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

func (r *Rect) X() int {
	return r.x
}

func (r *Rect) Y() int {
	return r.y
}

func (r *Rect) Width() int {
	return r.width
}

func (r *Rect) Height() int {
	return r.height
}
