package events

import "errors"

var ErrorHandlerAlreadyRegistered = errors.New("handler already registered")

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
				return ErrorHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) ClearHandlers() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}

func (ed *EventDispatcher) HasHandler(eventName string, handler EventHandlerInterface) bool {
	_, ok := ed.handlers[eventName]
	if ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) DispatchEvent(event EventInterface) error {
	_, ok := ed.handlers[event.GetName()]
	if ok {
		for _, h := range ed.handlers[event.GetName()] {
			err := h.HandleEvent(event)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
