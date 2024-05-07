package main

import (
	"fmt"
	"time"

	"github.com/datumforge/datum/pkg/events/emitter"
)

func main() {
	// Initialize a goroutine pool with 5 workers and a maximum capacity of 1000 tasks
	pool := emitter.NewPondPool(5, 1000)

	// Create a new emitter instance using the custom pool
	e := emitter.NewMemoryEmitter(emitter.WithPool(pool))

	// Define a listener that simulates a time-consuming task
	timeConsumingListener := func(evt emitter.Event) error {
		fmt.Printf("Processing event: %s with payload: %v\n", evt.Topic(), evt.Payload())
		// Simulate some work with a sleep
		time.Sleep(2 * time.Second)
		fmt.Printf("Finished processing event: %s\n", evt.Topic())
		return nil
	}

	// Subscribe the listener to a topic
	e.On("user.signup", timeConsumingListener)

	// Emit several events concurrently
	for i := 0; i < 10; i++ {
		go func(index int) {
			payload := fmt.Sprintf("User #%d", index)
			e.Emit("user.signup", payload)
		}(i)
	}

	// Wait for all events to be processed before shutting down
	time.Sleep(10 * time.Second)

	// Release the resources used by the pool
	pool.Release()
	fmt.Println("All events have been processed and the pool has been released.")
}
