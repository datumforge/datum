package marionette

import (
	"sort"
	"sync"
	"time"

	"go.uber.org/zap"
)

// Scheduler manages a list of future tasks and on or after the time that they are
// supposed to be scheduled, the scheduler sends the task on the out channel. This
// allows marionette to schedule tasks for the future or to retry tasks with backoff delays
// so that the workers are not overwhelmed by long running tasks.
//
// The scheduler is implemented with a sorted list of Futures (basically tasks arranged
// in time order) along with a single go routine that sleeps until the timestamp of the
// next task unless interrupted by a newly scheduled task. The goal of the scheduler is
// to minimize the number of go routines and CPU cycles in use so that higher priority
// work by the task manager or the main go routine is favored by the CPU. To that end
// the scheduler does not use a "ticker" clock checking if it should execute every
// second, preferring longer sleeps and interrupts instead.
type Scheduler struct {
	sync.RWMutex
	logger  *zap.SugaredLogger
	tasks   Futures
	out     chan<- Task
	add     chan *Future
	stop    chan struct{}
	running bool
}

// NewScheduler creates a new scheduler that can schedule task futures. The out channel is used to
// dispatch tasks at their scheduled time. If a task is sent on the out channel it means
// that the task should be executed as soon as possible. The scheduler makes no
// guarantees about exact timing of tasks scheduled except that the task will not be
// sent on the out channel before its scheduled time.
func NewScheduler(out chan<- Task, logger *zap.SugaredLogger) *Scheduler {
	return &Scheduler{
		out:     out,
		add:     make(chan *Future, 1),
		stop:    make(chan struct{}),
		tasks:   make(Futures, 0, minFuturesCapacity),
		running: false,
		logger:  logger,
	}
}

// Delay schedules the task to be run on or after the specified delay duration from now.
func (s *Scheduler) Delay(delay time.Duration, task Task) error {
	return s.Schedule(time.Now().Add(delay), task)
}

// Schedule a task to run on or after the specified timestamp. If the scheduler is
// running the task future is sent to the main channel loop, otherwise the tasks is
// simply inserted into the futures slice. Schedule blocks until the task is received
// by the main scheduler loop.
func (s *Scheduler) Schedule(at time.Time, task Task) error {
	future := &Future{Time: at, Task: task}
	if err := future.Validate(); err != nil {
		return err
	}

	s.Lock()
	if s.running {
		s.add <- future
	} else {
		s.tasks = s.tasks.Insert(future)
	}
	s.Unlock()

	return nil
}

// Start the scheduler in its own go routine or no-op if already started. If the
// specified wait group is not nil, it is marked as done when the scheduler is stopped.
func (s *Scheduler) Start(wg *sync.WaitGroup) {
	s.Lock()
	defer s.Unlock()

	if s.running {
		return
	}

	s.running = true

	if wg != nil {
		wg.Add(1)
	}

	go func() {
		s.run()

		if wg != nil {
			wg.Done()
		}
	}()
}

func (s *Scheduler) run() {
	s.logger.Info("scheduler running")

	// Schedule any tasks before or equal to now, ensuring that the next task in the
	// queue is in the future so that we can sleep until that timestamp.
	now := time.Now().In(time.UTC)
	s.schedule(now)

	// Start the scheduler loop
	for {
		// Create a delay timer based on the scheduled time of the next task so that
		// we're not waking periodically and checking. If there are no scheduled tasks
		// then just sleep for a day, newly schedule tasks will still be handled.
		//
		// NOTE: the timer needs to be created in the for block so a new timer is
		// allocated in each loop and old timers are not reused.
		var timer *time.Timer
		if len(s.tasks) == 0 || s.tasks[0].Time.IsZero() {
			timer = time.NewTimer(24 * time.Hour) // nolint: gomnd
		} else {
			timer = time.NewTimer(s.tasks[0].Time.Sub(now))
		}

		// Wait until either the timer goes off, a stop semaphore is sent, or a new
		// task has been scheduled before continuing the scheduler routine.
		select {
		case now = <-timer.C:
			s.schedule(now)

		case future := <-s.add:
			timer.Stop()

			now = time.Now().In(time.UTC)

			s.tasks = s.tasks.Insert(future)

		case <-s.stop:
			timer.Stop()
			s.logger.Info("scheduler stopped")

			return
		}
	}
}

// Sends all tasks that are before or equal to the specified timestamp on the out
// channel then resizes the tasks array to delete all futures that were sent.
func (s *Scheduler) schedule(at time.Time) {
	var sent int

scheduler:
	for _, future := range s.tasks {
		// Because all tasks are sorted if this task is after the timestamp, then we know
		// all tasks that follow it are also after the timestamp and we can stop.
		if future.Time.After(at) {
			break
		}

		// If the task is before or equal to the timestamp, send it on the out channel.
		// Perform a non-blocking send to ensure there are no scheduler deadlocks
		select {
		case s.out <- future.Task:
			sent++
		default:
			// If we couldn't send the task, stop trying to send and clean up the tasks
			// that were sent; the tasks will be resent on the next loop.
			break scheduler
		}
	}

	if sent > 0 {
		// If we sent tasks on the out channel, remove them from tasks and resize.
		s.tasks = s.tasks[sent:].Resize()
	}
}

// Stop the scheduler if it is running, otherwise a no-op. Note that stopping the
// scheduler does not close the out channel. When stopped, any futures that are still
// pending will not be executed, but if the scheduler is started again, they will remain
// as previously scheduled and sent on the same out channel.
func (s *Scheduler) Stop() {
	s.Lock()
	defer s.Unlock()

	if s.running {
		s.stop <- struct{}{}
		s.running = false
	}
}

func (s *Scheduler) IsRunning() bool {
	s.RLock()
	defer s.RUnlock()

	return s.running
}

// Future is a task/timestamp tuple that acts as a scheduler entry for running the task
// as close to the timestamp as possible without running it before the given time.
type Future struct {
	Time time.Time
	Task Task
}

func (f *Future) Validate() error {
	if f.Time.IsZero() {
		return ErrUnschedulable
	}

	return nil
}

// Futures implements the sort.Sort interface and ensures that the list of future tasks
// is maintained in sorted order so that tasks are scheduled correctly. This slice is
// also memory managed to ensure that it is garbage collected routinely and does not
// memory leak (e.g. using the Resize function to create a new slice and free the old).
type Futures []*Future

// The number of tasks we maintain in the futures slice to prevent allocations.
const minFuturesCapacity = 16

// Len implementation of the sort.Sort interface
func (f Futures) Len() int           { return len(f) }
func (f Futures) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f Futures) Less(i, j int) bool { return f[i].Time.Before(f[j].Time) }

// Insert a future into the slice of futures, growing the slice as necessary and
// returning it to replace the original slice (similar to append). Insert insures that
// the slice is maintained in sorted order and should be used instead of append.
func (f Futures) Insert(t *Future) Futures {
	f = append(f, nil) // extend the slice and make room

	// search for the position to insert the future then make room and add it
	i := sort.Search(len(f), func(i int) bool { return f[i] == nil || f[i].Time.After(t.Time) })
	copy(f[i+1:], f[i:])
	f[i] = t

	return f
}

// Resize the futures by copying the current futures into a new futures array, allowing
// the garbage collector to cleanup the previous slice and free up memory.
// See: https://forum.golangbridge.org/t/free-memory-of-slice/3713/2
func (f Futures) Resize() Futures {
	// Create new futures slice with size of old  futures but at least the specified cap
	var r Futures
	if len(f) < minFuturesCapacity {
		r = make(Futures, len(f), minFuturesCapacity)
	} else {
		r = make(Futures, len(f))
	}

	copy(r, f) // copy everything from f into r

	return r
}
