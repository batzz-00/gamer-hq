package events

import "fmt"

var handlerInstance *EventHandler

type Event struct {
	Name string
	Data interface{}
}

type EventHandler struct {
	Events      map[string]chan interface{}
	HookChannel chan Event
}

func newEventHandler() *EventHandler {
	return &EventHandler{
		Events:      make(map[string]chan interface{}),
		HookChannel: make(chan Event, 2),
	}
}

func AddEvent(e string) {
	Instance().Events[e] = make(chan interface{}, 2)
}

func Emit(event string, data interface{}) {
	if _, ok := Instance().Events[event]; ok {
		go func() {
			Instance().Events[event] <- data
		}()
		Instance().Hook(Event{Name: event, Data: data})
	} else {
		fmt.Println(fmt.Sprintf("Attempted to emit %s but not registered.", event))
	}
}

func (eh *EventHandler) Hook(event Event) {
	go func() {
		eh.HookChannel <- event
	}()
}

func Instance() *EventHandler {
	if handlerInstance == nil {
		handlerInstance = newEventHandler()
	}
	return handlerInstance
}
