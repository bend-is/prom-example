package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	listen = "172.17.0.1:6060"
	max    = 1_000_000
	system = "myapp"
)

type Metrics struct {
	calls        prometheus.Counter
	duration     prometheus.Histogram
	lastDuration prometheus.Gauge
}

var metrics = &Metrics{
	calls: promauto.NewCounter(prometheus.CounterOpts{Namespace: system, Name: "superjob_calls"}),
	duration: promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: system,
		Name:      "superjob_duration",
		Buckets:   []float64{.005, .01, .02, .03, .05, .1, .15, .95},
	}),
	lastDuration: promauto.NewGauge(prometheus.GaugeOpts{Namespace: system, Name: "superjob_last_duration"}),
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/superjob", superhandler)
	http.ListenAndServe(listen, nil)
}

func superhandler(w http.ResponseWriter, r *http.Request) {
	metrics.calls.Add(1)
	timer := prometheus.NewTimer(metrics.duration)
	defer func() {
		dur := timer.ObserveDuration()
		metrics.lastDuration.Set(dur.Seconds())
	}()

	size := rand.Int31n(max)
	slice := make([]int, size)
	for idx := range slice {
		slice[idx] = rand.Intn(max)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Len %d", len(slice))))
}
