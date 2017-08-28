package async

import(
	"log"
)

func StartDispatcher(handlerCount int) {
	dispatchQueue := make(chan chan Event, handlerCount)
	for i := 0; i < handlerCount; i++ {
		log.Printf("Starting asynchronous event handler %d", i+1)
		h := eventHandler{id: i+1}
		h.start(dispatchQueue)
	}
	// event processing loop
	go func() {
		for {
			select {
			case event := <- EventQueue:
				go func() {
					localEventQueue := <- dispatchQueue
					localEventQueue <- event
				}()
			}
		}
	}()
}
