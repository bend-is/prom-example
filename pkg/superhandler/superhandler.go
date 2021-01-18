package superhandler

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"net/http"
	"time"
)

const max = 1_000_000

type Handler struct {
	metrics *Metrics
}

func New(metrics *Metrics) *Handler {
	return &Handler{
		metrics: metrics,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	if h.metrics != nil {
		defer func() {
			delta := time.Since(t)

			h.metrics.Calls.Inc()
			h.metrics.Duration.Observe(delta.Seconds())
			h.metrics.DurationSum.Observe(delta.Seconds())
			h.metrics.LastDuration.Set(delta.Seconds())
			h.metrics.SummaryVec.With(prometheus.Labels{metricLabel: "handle_end"}).Observe(delta.Seconds())
		}()
	}

	size := rand.Int31n(max)
	slice := make([]int, size)

	if h.metrics != nil {
		h.metrics.SummaryVec.With(prometheus.Labels{metricLabel: "handle_preallocate"}).Observe(time.Since(t).Seconds())
	}

	for idx := range slice {
		slice[idx] = rand.Intn(max)
	}

	if h.metrics != nil {
		h.metrics.SummaryVec.With(prometheus.Labels{metricLabel: "handle_allocate"}).Observe(time.Since(t).Seconds())
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Len %d", len(slice))))
}
