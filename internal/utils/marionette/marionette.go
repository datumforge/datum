package marionette

import (
	"sync"
	"time"

	"go.uber.org/zap"
)

// TaskManager execute Tasks using a fixed number of workers that operate in their own
// go routines. The TaskManager also has a fixed task queue size, so that if there are
// more tasks added to the task manager than the queue size, back pressure is applied
type TaskManager struct {
	sync.RWMutex
	conf      Config
	logger    *zap.SugaredLogger
	scheduler *Scheduler
	wg        *sync.WaitGroup
	add       chan Task
	stop      chan struct{}
	running   bool
}

// New creates a new task manager with the specified configuration
func New(conf Config) *TaskManager {
	if conf.IsZero() {
		conf.Workers = 4
		conf.QueueSize = 64
		conf.ServerName = "marionette"
	}

	add := make(chan Task, conf.QueueSize)

	scheduler := NewScheduler(add, conf.Logger)

	return &TaskManager{
		conf:      conf,
		scheduler: scheduler,
		wg:        &sync.WaitGroup{},
		add:       add,
		stop:      make(chan struct{}, 1),
		running:   false,
		logger:    conf.Logger,
	}
}

// Queue a task to be executed asynchronously as soon as a worker is available; blocks if queue is full
func (tm *TaskManager) Queue(task Task, opts ...Option) error {
	handler := tm.WrapTask(task, opts...)

	tm.RLock()
	defer tm.RUnlock()

	if !tm.running {
		tm.logger.Warn(ErrTaskManagerStopped.Error())
		return ErrTaskManagerStopped
	}

	tm.add <- handler

	return nil
}

// Delay a task to be scheduled the specified duration from now
func (tm *TaskManager) Delay(delay time.Duration, task Task, opts ...Option) error {
	return tm.scheduler.Delay(delay, tm.WrapTask(task, opts...))
}

// Schedule a task to be executed at the specific timestamp
func (tm *TaskManager) Schedule(at time.Time, task Task, opts ...Option) error {
	return tm.scheduler.Schedule(at, tm.WrapTask(task, opts...))
}

// Start the task manager and scheduler in their own go routines (no-op if already started)
func (tm *TaskManager) Start() {
	tm.Lock()
	defer tm.Unlock()

	// Start the scheduler (also a no-op if already started)
	tm.scheduler.Start(tm.wg)

	if tm.running {
		return
	}

	tm.running = true
	go tm.run()
}

func (tm *TaskManager) run() {
	tm.wg.Add(1)
	defer tm.wg.Done()

	tm.logger.Info("task manager running")

	queue := make(chan *TaskHandler, tm.conf.QueueSize)

	for i := 0; i < tm.conf.Workers; i++ {
		tm.wg.Add(1)
		go worker(tm.wg, queue) // nolint: wsl
	}

	for {
		select {
		case task := <-tm.add:
			if handler, ok := task.(*TaskHandler); ok {
				queue <- handler
			} else {
				queue <- tm.WrapTask(task)
			}

		case <-tm.stop:
			close(queue)
			tm.logger.Info("task manager stopped")

			return
		}
	}
}

// worker function is a goroutine that executes tasks from the task queue. It receives
// tasks from the `tasks` channel and executes them by calling the `Exec` method on the `
// TaskHandler. This function runs in its own goroutine and is responsible for
// processing tasks concurrently
func worker(wg *sync.WaitGroup, tasks <-chan *TaskHandler) {
	defer wg.Done()

	for handler := range tasks {
		handler.Exec()
	}
}

// Stop stops the task manager and scheduler if running (otherwise a no-op). This method
// blocks until all pending tasks have been completed, however future scheduled tasks
// will likely be dropped and not scheduled for execution.
func (tm *TaskManager) Stop() {
	tm.Lock()

	// Stop the scheduler (also a no-op if already stopped)
	tm.scheduler.Stop()

	if tm.running {
		// Send the stop signal to the task manager
		tm.stop <- struct{}{}
		tm.running = false

		tm.Unlock()

		// Wait for all tasks to be completed and workers closed
		// TODO: write pending / future scheduled tasks somewhere?
		tm.wg.Wait()
	} else {
		tm.Unlock()
	}
}

// IsRunning checks if the taskmanager is running
func (tm *TaskManager) IsRunning() bool {
	tm.RLock()
	defer tm.RUnlock()

	return tm.running
}
