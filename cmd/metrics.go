package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	calls        prometheus.Counter
	duration     prometheus.Histogram
	durationSum  prometheus.Summary
	lastDuration prometheus.Gauge
}

func NewMetrics() *Metrics {
	return &Metrics{
		calls: prometheus.NewCounter(prometheus.CounterOpts{Namespace: system, Name: "superjob_calls"}),
		duration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: system,
			Name:      "superjob_duration",
			Buckets:   []float64{.005, .01, .02, .03, .05, .1, .15, .95},
		}),
		durationSum: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace:  system,
			Name:       "superjob_summary",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}),
		lastDuration: prometheus.NewGauge(prometheus.GaugeOpts{Namespace: system, Name: "superjob_last_duration"}),
	}
}
