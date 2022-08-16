package events

import "github.com/filipgorny/editor/internal/types"

type KeyEvent struct {
	id  types.ID
	key string
}

func NewKeyEvent(key string) *KeyEvent {
	return &KeyEvent{
		id:  types.NewId(),
		key: key,
	}
}

func (e *KeyEvent) Id() types.ID {
	return e.id
}

func (e *KeyEvent) Type() string {
	return "key"
}

func (e *KeyEvent) Key() string {
	return e.key
}
