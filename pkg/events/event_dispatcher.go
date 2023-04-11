package events

import "errors"

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) RegisterHandler(eventName string, handler EventHandlerInterface) error {
	_, ok := ed.handlers[eventName]
	if ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return errors.New("handler already registered")
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}
