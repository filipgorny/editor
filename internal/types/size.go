package types

type Size struct {
	width  int
	height int
}

func NewSize(width int, height int) *Size {
	return &Size{
		width:  width,
		height: height,
	}
}

func (s *Size) Width() int {
	return s.width
}

func (s *Size) Height() int {
	return s.height
}
