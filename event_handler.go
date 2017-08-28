package async

import (
	"log"
)

type eventHandler struct {id int}

func (h *eventHandler) start(dispatchQueue chan chan Event) {
	localEventQueue := make(chan Event)
	go func() {
		for {
			dispatchQueue <- localEventQueue
			select {
			case event := <- localEventQueue:
				log.Printf("Asynchronous event handler %d is processing an event.\n", h.id)
				event.Process()
			}
		}
	}()
}


