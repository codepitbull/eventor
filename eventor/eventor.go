package eventor

import (
	"errors"
	"fmt"
	"time"
)

type Step string

type EventFunc func() (Step, time.Duration, error)

type EventProcessor struct {
	start Step
	eventHandlers map[Step]EventFunc
}


func NewEventProcessor(start Step, eventHandlers map[Step]EventFunc) EventProcessor {
	return EventProcessor{
		start: start,
		eventHandlers: eventHandlers,
	}
}

func (e *EventProcessor) Run() (Step, error) {
	if e.start == "" {
		return "", errors.New("no start step defined")
	}
	return e.RunFrom(e.start)
}

func (e *EventProcessor) RunFrom(step Step) (Step, error) {
	fun, ok := e.eventHandlers[step]
	if !ok {
		return "", fmt.Errorf("unknown step: %s", step)
	}
	next, duration, err := fun()
	if err != nil {
		return "", err
	}
	if next == "" {
		return "", nil
	}
	if duration > 0 {
		return next, nil
	}
	return e.RunFrom(next)
}
