package main

import (
	"fmt"
	"time"
)

type (
	Event struct {
		Data  int64
		Count int64
	}

	Observer interface {
		OnNotify(event Event)
	}

	Notifier interface {
		Register(observer Observer)
		Deregister(observer Observer)
		Notify(event Event)
	}
)

type (
	eventObserver struct {
		id int
	}

	eventNotifier struct {
		observers map[Observer]struct{}
	}
)

func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("*** Observer %d received: %d count: %d\n", o.id, e.Data, e.Count)
}

func (e *eventNotifier) Register(o Observer) {
	e.observers[o] = struct{}{}
}

func (e *eventNotifier) Deregister(o Observer) {
	delete(e.observers, o)
}

func (p *eventNotifier) Notify(e Event) {
	for o := range p.observers {
		o.OnNotify(e)
	}
}

func main() {
	// Initialize a new Notifier
	n := eventNotifier{
		observers: map[Observer]struct{}{},
	}

	// Register a couple of observers
	n.Register(&eventObserver{id: 1})
	n.Register(&eventObserver{id: 2})

	// simple loop publishing the current Unix timestamp to observers
	stop := time.NewTicker(10 * time.Second).C
	tick := time.NewTicker(time.Millisecond).C
	var i int64 = 0
	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			i++
			n.Notify(Event{Data: t.UnixNano(), Count: i})
		}
	}
}
