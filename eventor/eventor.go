package eventor

import (
	"time"
)

type EventFunc func() (*time.Duration, error)

type EventProcessor struct {
	eventHandlers []EventFunc
}


func NewEventProcessor(eventHandlers... EventFunc) EventProcessor {
	return EventProcessor{
		eventHandlers: eventHandlers,
	}
}

func (f *EventProcessor) Run() error {
	for _, handler := range f.eventHandlers {
		if _, err := handler(); err != nil {
			return err
		}
	}
	return nil
}
