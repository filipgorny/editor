package bus

type Bus struct {
	listeners map[string]Listener
}

func NewBus() *Bus {
	return &Bus{
		listeners: make(map[string]Listener),
	}
}
