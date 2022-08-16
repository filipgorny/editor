package sensor

import "github.com/filipgorny/editor/internal/editor/events"

type SensorEventCallback func(e events.Event)

type Sensor interface {
	Run(callback SensorEventCallback)
}
