package main

type EventStore struct {
	events   []*Event
	handlers []*EventHandler
}

func (store *EventStore) AddEvent(event *Event) *AddEventError {
	if _, hasEvent := store.getEvent(event); hasEvent {
		return &AddEventError{}
	}

	store.events = append(store.events, event)

	store.notifyHandlers(event)

	return nil
}

func (store *EventStore) RemoveEvent(event *Event) *RemoveEventError {
	index, hasRemovedEvent := store.getEvent(event)
	if !hasRemovedEvent {
		return &RemoveEventError{}
	}

	eventCount := len(store.events)
	var newEvents []*Event

	switch {
	case index == 0:
		newEvents = store.events[1:eventCount]
	case index == eventCount-1:
		newEvents = store.events[:eventCount-1]
	default:
		newEvents = store.events[:index]
		newEvents = append(newEvents, store.events[index+1:]...)
	}

	store.events = newEvents

	return nil
}

func (store *EventStore) UpdateEvent(oldEvent, newEvent *Event) *UpdateEventError {
	index, hasOldEvent := store.getEvent(oldEvent)
	if !hasOldEvent {
		return &UpdateEventError{}
	}

	store.events[index] = newEvent

	store.notifyHandlers(newEvent)

	return nil
}

func (store *EventStore) AddHandler(handler *EventHandler) *AddHandlerError {
	if _, hasHandler := store.getHandler(handler); hasHandler {
		return &AddHandlerError{}
	}

	store.handlers = append(store.handlers, handler)

	return nil
}

func (store *EventStore) RemoveHandler(handler *EventHandler) *RemoveHandlerError {
	index, hasRemovedHandler := store.getHandler(handler)
	if !hasRemovedHandler {
		return &RemoveHandlerError{}
	}

	handlerCount := len(store.handlers)
	var newHandlers []*EventHandler

	switch {
	case index == 0:
		newHandlers = store.handlers[1:handlerCount]
	case index == handlerCount-1:
		newHandlers = store.handlers[:handlerCount-1]
	default:
		newHandlers = store.handlers[:index]
		newHandlers = append(newHandlers, store.handlers[index+1:]...)
	}

	store.handlers = newHandlers

	return nil
}

func (store *EventStore) getEvent(searchedEvent *Event) (int, bool) {
	for i, event := range store.events {
		if event == searchedEvent {
			return i, true
		}
	}

	return 0, false
}

func (store *EventStore) getHandler(searchedHandler *EventHandler) (int, bool) {
	for i, handler := range store.handlers {
		if handler == searchedHandler {
			return i, true
		}
	}

	return 0, false
}

func (store *EventStore) notifyHandlers(event *Event) {
	for _, handler := range store.handlers {
		(*handler)(event)
	}
}
