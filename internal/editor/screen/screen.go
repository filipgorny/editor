package screen

import (
	"github.com/filipgorny/editor/internal/editor/workspace"
	"github.com/filipgorny/editor/internal/types"
)

type Screen struct {
	size      types.Size
	workspace *workspace.Workspace
}

func NewScreen(size types.Size, workspace *workspace.Workspace) *Screen {
	return &Screen{
		size:      size,
		workspace: workspace,
	}
}

func (s *Screen) Workspace() *workspace.Workspace {
	return s.workspace
}
