package soiree

import "github.com/prometheus/client_golang/prometheus"

func (p *PondPool) NewStatsCollector() {
	prometheus.MustRegister(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "p_workers_running",
			Help: "Number of running worker goroutines",
		},
		func() float64 {
			return float64(p.Running())
		}))
	prometheus.MustRegister(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "p_workers_idle",
			Help: "Number of idle worker goroutines",
		},
		func() float64 {
			return float64(p.IdleWorkers())
		}))

	// Task metrics
	prometheus.MustRegister(prometheus.NewCounterFunc(
		prometheus.CounterOpts{
			Name: "p_tasks_submitted_total",
			Help: "Number of tasks submitted",
		},
		func() float64 {
			return float64(p.SubmittedTasks())
		}))
	prometheus.MustRegister(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "p_tasks_waiting_total",
			Help: "Number of tasks waiting in the queue",
		},
		func() float64 {
			return float64(p.WaitingTasks())
		}))
	prometheus.MustRegister(prometheus.NewCounterFunc(
		prometheus.CounterOpts{
			Name: "p_tasks_successful_total",
			Help: "Number of tasks that completed successfully",
		},
		func() float64 {
			return float64(p.SuccessfulTasks())
		}))
	prometheus.MustRegister(prometheus.NewCounterFunc(
		prometheus.CounterOpts{
			Name: "p_tasks_failed_total",
			Help: "Number of tasks that completed with panic",
		},
		func() float64 {
			return float64(p.FailedTasks())
		}))
	prometheus.MustRegister(prometheus.NewCounterFunc(
		prometheus.CounterOpts{
			Name: "p_tasks_completed_total",
			Help: "Number of tasks that completed either successfully or with panic",
		},
		func() float64 {
			return float64(p.CompletedTasks())
		}))
}
