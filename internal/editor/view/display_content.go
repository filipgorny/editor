package view

import "github.com/filipgorny/editor/internal/types"

type DisplayContent struct {
	lines map[int]string
	size  *types.Size
}

func NewDisplayContent(width int, height int) *DisplayContent {
	return &DisplayContent{
		lines: make(map[int]string),
		size:  types.NewSize(width, height),
	}
}

func (d *DisplayContent) Lines() map[int]string {
	return d.lines
}

func (d *DisplayContent) Write(i int, s string) {
	d.lines[i] = s
}

func (d *DisplayContent) Get(x int, y int) rune {
	return rune(d.lines[y][x])
}
