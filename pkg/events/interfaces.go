package events

import "time"

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	HandleEvent(event EventInterface) error
}

type EventDispatcherInterface interface {
	RegisterHandler(eventName string, handler EventHandlerInterface) error
	DispatchEvent(event EventInterface) error
	RemoveHandler(eventName string, handler EventHandlerInterface) error
	HasHandler(eventName string, handler EventHandlerInterface) bool
	ClearHandlers(eventName string) error
}