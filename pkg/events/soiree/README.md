# Soiree

Soiree, a fancy event affair, or, event - is a library indendied to simplify event management. The goal is a dead-simple interface for event subscription and handling, using [pond](https://github.com/alitto/pond) for performance and wrapping event management with thread-safe interactions.

## Overview

Functionally, Soiree is intended to provide:

- **In-Memory management**: Host and manage events internally without external dependencies or libraries
- **Listener prioritization**: Controls for invocation order
- **Concurrent Processing**: Utilize pooled goroutines for handling events in parallel and re-releasing resources, with thread safety
- **Configurable Subscriptions**: Leverages basic pattern matching for event subscriptions
- **General Re-use**: Configure with custom handlers for errors, IDs, and panics panics

### But why?

In modern software architectures, microservices vs. monoliths is a false dichotomy; the optimum is usually somewhere in the middle. With the Datum service, we're pretty intentionally sticking with a "monolith" in the sense we are producing a single docker image from this codebase, but the often overlooked aspect of these architectures is the context under which the service is started, and run. If you assume that the connectivity from your client is created in a homogeneous fashion, ex:

┌──────────────┐        .───────.         ┌─────────────────┐
│              │       ╱         ╲        │                 │
│    Client    │──────▶   proxy   ────────▶  Datum Service  │
│              │       `.       ,'        │                 │
└──────────────┘         `─────'          └─────────────────┘

Then all instances of the datum service will be required to perform things like authorizations validation, session issuance, etc. The validity of these actions is managed with external state machines, such as Redis.

                                                                   ┌────────────────┐
┌──────────────┐        .───────.         ┌─────────────────┐      │                │
│              │       ╱         ╲        │                 │      │     Redis,     │
│    Client    │──────▶   proxy   ────────▶  Datum Service  ├─────▶│   PostgreSQL   │
│              │       `.       ,'        │                 │      │                │
└──────────────┘         `─────'          └─────────────────┘      └────────────────┘

We do this because we want to be able to run many instances of the Datum service, for things such as canary, rollouts, etc.

                                          ┌─────────────────┐
                                          │                 │
                                     ┌───▶│  Datum Service  ├──┐
                                     │    │                 │  │
                                     │    └─────────────────┘  │
                                     │                         │   ┌────────────────┐
┌──────────────┐        .───────.    │    ┌─────────────────┐  │   │                │
│              │       ╱         ╲   │    │                 │  │   │     Redis,     │
│    Client    │──────▶   proxy   ───┼────▶  Datum Service  ├──┴┬─▶▶   PostgreSQL   │
│              │       `.       ,'   │    │                 │   │  │                │
└──────────────┘         `─────'     │    └─────────────────┘   │  └────────────────┘
                                     │                          │
                                     │    ┌─────────────────┐   │
                                     │    │                 │   │
                                     └───▶│  Datum Service  │───┘
                                          │                 │
                                          └─────────────────┘

Now, where things start to get more fun is when you layer in the desire to perform I/O operations either managed by us, or externally (e.g. S3), as well as connect to external data stores (e.g. Turso).

                                             ┌──────────────┐
                                             │              │
                                             │      S3      │
                                             │              │           ┌───────────────┐
                                             └───────▲──────┘    ┌─────▶│ Outbound HTTP │
                                                     │           │      │(e.g. webhooks)│
                                                     │           │      └───────────────┘
                                                     │           │
                                                   ┌─┘           │
                                                   │             │
                   ┌ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┼ ─ ─ ─ ─ ─ ─ ┼ ─ ─ ─ ┬─────────────┐
                                                   │             │       │ Stuff under │
                   │                               │             │       │ our control │
                                          ┌────────┴────────┬────┘       └─────────────┘
                   │                      │                 │                          │
                                     ┌───▶│  Datum Service  ├──┐
                   │                 │    │                 │  │                       │
                                     │    └─────────────────┘  │
                   │                 │                         │   ┌────────────────┐  │
┌──────────────┐        .───────.    │    ┌─────────────────┐  │   │                │
│              │   │   ╱         ╲   │    │                 │  │   │     Redis,     │  │
│    Client    │──────▶   proxy   ───┼────▶  Datum Service  ├──┴┬─▶▶   PostgreSQL   │
│              │   │   `.       ,'   │    │                 │   │  │                │  │
└──────────────┘         `─────'     │    └─────────────────┘   │  └────────────────┘
                   │                 │                          │                      │
                                     │    ┌─────────────────┐   │
                   │                 │    │                 │   │                      │
                                     └───▶│  Datum Service  │───┘
                   │                      │                 │                          │
                                          └────────┬────────┴─────────────┐
                   │                               │                      │            │
                    ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─│─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┼ ─ ─ ─ ─ ─ ─
                                                   │                      │
                                                   └─────┐                │
                                                         │                │
                                                         │                │               ┌──────────────┐
                                                         │                │               │ Other future │
                                                         ▼                └──────────────▶│  ridiculous  │
                                                     ┌───────────┐                        │    ideas     │
                                                     │   Turso   │                        └──────────────┘
                                                     └───────────┘


Adding "events" is really not as simple as having something like Kafka or Redis, either. In a "traditional" event architecture you will usually have something like this:

                             ┌────────────────────────┐
                             │   Traditional Event    │
                             │     Architectures      │
                             └────────────────────────┘

┌────────────────┐
│                │                                                                ┌────────────────┐
│   Publisher    │────┐                                                           │                │
│                │    │                                                   ┌───────▶   Subscriber   │
└────────────────┘    │                                                   │       │                │
                 ┌────┴──────┐                                      ┌───────────┐ └────────────────┘
                 │ Message 1 │                                      │ Messages  │
                 └────┬──────┘              .───────────────.       └───────────┘
┌────────────────┐    │                 _.─'                 `──.         │       ┌────────────────┐
│                │   ┌┴──────────┐     ╱    Pub / Sub Service    ╲        │       │                │
│   Publisher    ├───┤ Message 3 ├────▶                           ───────┬┴───────▶   Subscriber   │
│                │   └─────┬─────┘     `.        Channel        ,'       │        │                │
└────────────────┘         │             `──.               _.─'         │        └────────────────┘
                           │                 `─────────────'             │
                           │                                             │
┌────────────────┐    ┌────┴──────┐                                      │        ┌────────────────┐
│                │    │ Message 2 │                                      │        │                │
│   Publisher    ├────┴───────────┘                                      └───────▶│   Subscriber   │
│                │                                                                │                │
└────────────────┘                                                                └────────────────┘


But this pattern assumes we can have discreet, secondary code bases consuming off an event broker channel. as well as consistent interfaces, data types, etc., for all those events. Pretty complicated. Creating more external dependencies for runtime is expensive in more ways than dollars, as well. What problem are we trying to solve for, then?

If Datum as a service, which looks a bit like what you'd traditionally consider a "monolith", needs to perform all kinds of workload actions such as writing a file to a bucket, committing a SQL transaction, sending an http request, we need bounded patterns. Which is to say, we need to control resource contention in our runtime as we don't want someone's regular HTTP request to be blocked by the fact someone else requested a bulk upload to S3. This means creating rough groupings of workload `types` and bounding them so that you can monitor and control the behaviors  and lumpiness of the variances with the workload types.

Additionally, unless we want to be in a living hell of reconciliation with regards to Event emission / management / storage, we're better served by creating in-memory constructs to execute on various action types. Example: we want to provide to our users the ability to send emails based on events which occur related to their use of Datum. If you were to offload this to an intermediary, such as datum -> kafka -> emailsender, in the event that the datum or kafka service fails, you're left asking the question if the email has been sent or not (e.g. has the event been processed).

## How many goroutines can / should I have?

A single go-routine currently uses a minimum stack size of 2KB. It is likely that your actual code will also allocate some additional memory on the heap (for e.g. JSON serialization or similar) for each goroutine. This would mean 1M go-routines could easily require 2-4 GB or RAM (should be ok for an average environment)

Most OS will limit the number of open connections in various ways. For TCP/IP there is usually a limit of open ports per interface. This is about 28K on many modern systems. There is usually an additional limit per process (e.g. ulimit for number of open file-descriptors) which will by default be around 1000. So without changing the OS configuration you will have a maximum of 1000 concurrent connections on Linux.

So depending on the system you should probably not create more than 1000 goroutines, because they might start failing with “maximum num of file descriptors reached” or even drop packets. If you increase the limits you are still bound by the 28K connections from a single IP address.

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/datumforge/datum/pkg/events/soiree"
)

func main() {
	e := soiree.NewWhisper()
	e.On("user.created", func(evt soiree.Event) error {
		fmt.Println("Event received:", evt.Topic())
		return nil
	})
	e.Emit("user.created", "Matty Ice")
}
```

## Configuration

Customize Soiree with the provided options:

```go
e := soiree.NewWhisper(
	soiree.WithErrorHandler(customErrorHandler),
	soiree.WithIDGenerator(customIDGenerator),
	// More options...
)
```

### Options

| Option                                         | Description                                                  |
|------------------------------------------------|--------------------------------------------------------------|
| `WithPool(pool soiree.Pool)`                  | Assign a goroutine pool for concurrent event handling       |
| `WithErrorHandler(handler func(soiree.Event, error) error)` | Set a custom error handler for the soiree that receives an event and an error |
| `WithIDGenerator(generator func() string)`     | Define a function for generating unique listener IDs        |
| `WithPanicHandler(handler func(interface{}))`  | Implement a panic recovery strategy                        |

## Pattern-matched event Subscription

Pattern-match event topics with wildcards:

- `*` - Matches a single segment
- `**` - Matches multiple segments

### Example:

```go
e := soiree.NewWhisper()
e.On("user.*", userEventListener)
e.On("invoice.**", orderEventListener)
e.On("**.completed", completionEventListener)
```

### Another Example

```go
e := soiree.NewWhisper()
e.On("user.*", func(evt soiree.Event) error {
	fmt.Printf("Event: %s, Payload: %+v\n", evt.Topic(), evt.Payload())
	return nil
})
e.Emit("user.signup", "Funky Sarah")
// Use synchronization instead of sleep in production.
```

## Aborting Event Propagation

Stop event propagation using `SetAborted`:

```go
e := soiree.NewWhisper()
e.On("invoice.processed", func(evt soiree.Event) error {
	if /* condition fails */ false {
		evt.SetAborted(true)
	}
	return nil
}, soiree.WithPriority(soiree.High))
e.On("invoice.processed", func(evt soiree.Event) error {
	// This will not run if the event is aborted
	return nil
}, soiree.WithPriority(soiree.Low))
e.Emit("invoice.processed", "Order data")
```

Abort event handling early based on custom logic

## Examples

- [Managing Concurrency](#managing-concurrency-with-withpool)
- [Custom Error Handling](#custom-error-handling-with-witherrorhandler)
- [Listener Prioritization](#prioritizing-listeners-with-withpriority)
- [ID Generation](#generating-unique-ids-with-withidgenerator)
- [Panic Recovery](#handling-panics-gracefully-with-withpanichandler)

### Managing Concurrency with `WithPool`

Delegate concurrency management to a custom goroutine pool using the `WithPool` option:

```go
package main

import (
	"github.com/datumforge/datum/pkg/events/soiree"
	"github.com/alitto/pond"
)

func main() {
	// Initialize a goroutine pool
	pool := soiree.NewPondPool(10, 1000) // 10 workers, queue size 1000

	// Set up the soiree with this pool
	e := soiree.NewWhisper(soiree.WithPool(pool))

	// Your soiree is now ready to handle events using the pool
}
```

### Custom Error Handling with `WithErrorHandler`

Enhance error visibility by defining a custom error handler:

```go
package main

import (
	"log"
	"github.com/datumforge/datum/pkg/events/soiree"
)

func main() {
	// Define a custom error handler that logs the event and the error
	customErrorHandler := func(event soiree.Event, err error) error {
		log.Printf("Error encountered during event '%s': %v, with payload: %v", event.Topic(), err, event.Payload())
		return nil  // Returning nil to indicate that the error has been handled
	}

	// Apply the custom error handler to the soiree
	e := soiree.NewWhisper(soiree.WithErrorHandler(customErrorHandler))

	// Your soiree will now log detailed errors encountered during event handling
}
```

### Prioritizing Listeners with `WithPriority`

Control the invocation order of event listeners:

```go
package main

import (
	"fmt"
	"github.com/datumforge/datum/pkg/events/soiree"
)

func main() {
	// Set up the soiree
	e := soiree.NewWhisper()

	// Define listeners with varying priorities
	normalPriorityListener := func(e soiree.Event) error {
		fmt.Println("Normal priority: Received", e.Topic())
		return nil
	}

	highPriorityListener := func(e soiree.Event) error {
		fmt.Println("High priority: Received", e.Topic())
		return nil
	}

	// Subscribe listeners with specified priorities
	e.On("user.created", normalPriorityListener) // Default is normal priority
	e.On("user.created", highPriorityListener, soiree.WithPriority(soiree.High))

	// Emit an event and observe the order of listener notification
	e.Emit("user.created", "User signup event")
}
```

Listeners with higher priority are notified first when an event occurs.

### Generating Unique IDs with `WithIDGenerator`

Implement custom ID generation for listener tracking:

```go
package main

import (
	"github.com/google/uuid"
	"github.com/datumforge/datum/pkg/events/soiree"
)

func main() {
	// Custom ID generator using UUID v4
	uuidGenerator := func() string {
		return uuid.NewString()
	}

	// Initialize the soiree with the UUID generator
	e := soiree.NewWhisper(soiree.WithIDGenerator(uuidGenerator))

	// Listeners will now be registered with a unique UUID
}
```

Listeners are now identified by a UUID vs. the standard ULID generated by Datum.

### Handling Panics Gracefully with `WithPanicHandler`

Safeguard your application from unexpected panics during event handling:

```go
package main

import (
	"log"
	"github.com/datumforge/datum/pkg/events/soiree"
)

func main() {
	// Define a panic handler that logs the occurrence
	logPanicHandler := func(p interface{}) {
		log.Printf("Panic recovered: %v", p)
		// Insert additional logic for panic recovery here
	}

	// Equip the soiree with the panic handler
	e := soiree.NewWhisper(soiree.WithPanicHandler(logPanicHandler))

	// Your soiree is now more resilient to panics
}
```

This handler ensures that panics are logged and managed without creating a service interruption for users.

