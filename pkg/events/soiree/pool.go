package soiree

import (
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

// Running returns the number of running workers in the pool
func (p *PondPool) Running() int {
	return p.pool.RunningWorkers()
}

// Release stops all workers in the pool and waits for them to finish
func (p *PondPool) Release() {
	p.pool.StopAndWait()
}
