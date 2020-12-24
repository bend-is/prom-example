package superhandler

import (
	"fmt"
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
	if h.metrics != nil {
		t := time.Now()

		defer func() {
			delta := time.Since(t)

			h.metrics.Calls.Inc()
			h.metrics.Duration.Observe(delta.Seconds())
			h.metrics.DurationSum.Observe(delta.Seconds())
			h.metrics.LastDuration.Set(delta.Seconds())
		}()
	}

	size := rand.Int31n(max)
	slice := make([]int, size)
	for idx := range slice {
		slice[idx] = rand.Intn(max)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Len %d", len(slice))))
}
