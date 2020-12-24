package main

import (
	"net/http"

	"github.com/korjavin/prom-example/pkg/superhandler"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	listen = "172.17.0.1:6060"
	system = "myapp"
	name   = "superjob"
)

func main() {
	metrics := superhandler.NewMetrics(system, name)
	prometheus.MustRegister(metrics.Calls, metrics.Duration, metrics.DurationSum, metrics.LastDuration)

	h := superhandler.New(metrics)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/superjob", h.Handle)
	http.ListenAndServe(listen, nil)
}
