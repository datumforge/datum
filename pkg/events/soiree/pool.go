package soiree

import (
	"time"

	"github.com/alitto/pond"
)

// Pool is an interface for a worker pool
type Pool interface {
	Submit(task func())
	Running() int
	Release()
}

// PondPool is a worker pool implementation using the pond library
type PondPool struct {
	pool *pond.WorkerPool
}

// NewPondPool creates a new instance of PondPool with the passed options
func NewPondPool(maxWorkers, maxCapacity int, options ...pond.Option) *PondPool {
	return &PondPool{
		pool: pond.New(maxWorkers, maxCapacity, options...),
	}
}

// Submit submits a task to the worker pool
func (p *PondPool) Submit(task func()) {
	p.pool.Submit(task)
}

// SubmitandWait submits a task to the worker pool and waits for it to finish
func (p *PondPool) SubmitAndWait(task func()) {
	p.pool.SubmitAndWait(task)
}

// SubmitBefore submits a task to the worker pool before a specified task
func (p *PondPool) SubmitBefore(task func(), deadline time.Duration) {
	p.pool.SubmitBefore(task, deadline)
}

// Running returns the number of running workers in the pool
func (p *PondPool) Running() int {
	return p.pool.RunningWorkers()
}

// Release stops all workers in the pool and waits for them to finish
func (p *PondPool) Release() {
	p.pool.StopAndWait()
}

// ReleaseWithDeadline stops this pool and waits until either all tasks in the queue are completed
// or the given deadline is reached, whichever comes first
func (p *PondPool) ReleaseWithDeadline(deadline time.Duration) {
	p.pool.StopAndWaitFor(deadline)
}

// Stop scauses this pool to stop accepting new tasks and signals all workers to exit
// Tasks being executed by workers will continue until completion (unless the process is terminated)
// Tasks in the queue will not be executed (so will drop any buffered tasks - ideally use Release or ReleaseWithDeadline)
func (p *PondPool) Stop() {
	p.pool.Stop()
}

// IdleWorkers returns the number of idle workers in the pool
func (p *PondPool) IdleWorkers() int {
	return p.pool.IdleWorkers()
}

// SubmittedTasks returns the number of tasks submitted to the pool
func (p *PondPool) SubmittedTasks() int {
	return int(p.pool.SubmittedTasks())
}

// WaitingTasks returns the number of tasks waiting in the pool
func (p *PondPool) WaitingTasks() int {
	return int(p.pool.WaitingTasks())
}

// SuccessfulTasks returns the number of tasks that completed successfully
func (p *PondPool) SuccessfulTasks() int {
	return int(p.pool.SuccessfulTasks())
}

// FailedTasks returns the number of tasks that completed with a panic
func (p *PondPool) FailedTasks() int {
	return int(p.pool.FailedTasks())
}

// CompletedTasks returns the number of tasks that completed either successfully or with a panic
func (p *PondPool) CompletedTasks() int {
	return int(p.pool.CompletedTasks())
}

// StopAndWaitFor stops this pool and waits until either all tasks in the queue are completed
// or the given deadline is reached, whichever comes first
func (p *PondPool) StopAndWaitFor(deadline time.Duration) {
	p.pool.StopAndWaitFor(deadline)
}