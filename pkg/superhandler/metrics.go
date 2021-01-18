package superhandler

import (
	"github.com/prometheus/client_golang/prometheus"
)

const metricLabel = "operation"

type Metrics struct {
	Calls        prometheus.Counter
	Duration     prometheus.Histogram
	DurationSum  prometheus.Summary
	LastDuration prometheus.Gauge
	SummaryVec   *prometheus.SummaryVec
}

func NewMetrics(namespace, name string) *Metrics {
	return &Metrics{
		Calls: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      name,
			Subsystem: "calls",
		}),
		Duration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      name,
			Subsystem: "duration",
			Buckets:   []float64{.005, .01, .02, .03, .05, .1, .15, .95},
		}),
		DurationSum: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace:  namespace,
			Name:       name,
			Subsystem:  "summary",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}),
		LastDuration: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      name,
			Subsystem: "last_duration",
		}),
		SummaryVec: prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Namespace:  namespace,
			Name:       name,
			Subsystem:  "duration_summary_vec",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}, []string{metricLabel}),
	}
}
