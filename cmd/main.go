package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	listen = "172.17.0.1:6060"
	max    = 1_000_000
	system = "myapp"
)

type Handler struct {
	metrics *Metrics
}

func main() {
	h := &Handler{metrics: NewMetrics()}
	prometheus.MustRegister(h.metrics.calls, h.metrics.duration, h.metrics.durationSum, h.metrics.lastDuration)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/superjob", h.superhandler)
	http.ListenAndServe(listen, nil)
}

func (h *Handler) superhandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	defer func() {
		delta := time.Since(t)

		h.metrics.calls.Inc()
		h.metrics.duration.Observe(delta.Seconds())
		h.metrics.durationSum.Observe(delta.Seconds())
		h.metrics.lastDuration.Set(delta.Seconds())
	}()

	size := rand.Int31n(max)
	slice := make([]int, size)
	for idx := range slice {
		slice[idx] = rand.Intn(max)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Len %d", len(slice))))
}
