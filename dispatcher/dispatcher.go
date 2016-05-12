package dispatcher

import "sort"

// Listener is a callback that is triggered on an event
type Listener func(Event)

type listeners []Listener

type listenersList map[int]listeners

type registry map[string]listenersList

// Dispatcher registers listeners as callbacks that should react to an event
type Dispatcher struct {
	listeners registry
}

// NewDispatcher creates and returns a new dispatcher instance
func NewDispatcher() Dispatcher {
	dispatcher := Dispatcher{}
	dispatcher.listeners = make(registry)
	return dispatcher
}

// Return a slice of listeners ordered by priority
func (l listenersList) GetPrioritisedListeners() listeners {
	i := 0
	priorities := make([]int, len(l))

	for priority := range l {
		priorities[i] = priority
		i++
	}

	sort.Ints(priorities)

	listeners := *new(listeners)

	for x := 0; x < len(priorities); x++ {
		listeners = append(listeners, l[priorities[x]]...)
	}

	return listeners
}

// Listen registers a listener to be triggered in an event
func (d *Dispatcher) Listen(eventName string, listener Listener, priority int) {
	if len(d.listeners[eventName]) == 0 {
		d.listeners[eventName] = make(listenersList, 0)
	}

	d.listeners[eventName][priority] = append(d.listeners[eventName][priority], listener)
}

// Dispatch triggers an event by name, passing an event to the listener
func (d *Dispatcher) Dispatch(eventName string, event Event) {
	listeners := d.listeners[eventName].GetPrioritisedListeners()

	for _, listener := range listeners {
		listener(event)
	}
}
