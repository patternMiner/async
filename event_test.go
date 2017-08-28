package async

import (
	"testing"
	"sync"
)

type testEvent struct{feedback chan string}

func (e testEvent) Process() {
	e.feedback <- "Hello"
}

var wg sync.WaitGroup

func TestEventDispatcher(t *testing.T) {
	StartDispatcher(4)
	feedback := make(chan string, 16)
	responseCount := 0

	wg.Add(16)
	for i := 0; i < 16; i++ {
		EventQueue <- testEvent{feedback: feedback}
	}

	go func() {
		for {
			select {
			case resp := <- feedback:
				if resp == "Hello" {
					responseCount++
				}
				wg.Done()
			}
		}
	}()

	wg.Wait()

	if responseCount != 16 {
		t.Errorf("Expected %d responseCount, but got %d", 16, responseCount)
	}
}