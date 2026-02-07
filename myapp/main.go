package main

import (
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metrics struct {
	opsProcessed prometheus.Counter
}

func newMetrics(reg prometheus.Registerer, event string, isSuccess bool) *metrics {
	reg = prometheus.WrapRegistererWith(prometheus.Labels{"event": event, "isSuccess": strconv.FormatBool(isSuccess)}, reg)
	m := &metrics{
		opsProcessed: promauto.With(reg).NewCounter(prometheus.CounterOpts{
			Name: "myapp_processed_ops_total",
		}),
	}
	return m
}

func recordMetrics1s(m *metrics) {
	go func() {
		for {
			m.opsProcessed.Inc()
			time.Sleep(time.Second)
		}
	}()
}

func recordMetrics2s(m *metrics) {
	go func() {
		for {
			m.opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func recordMetricsRs(m *metrics) {
	go func() {
		for {
			r := rand.IntN(13)
			m.opsProcessed.Inc()
			time.Sleep(time.Duration(r) * time.Second)
		}
	}()
}

func main() {
	reg := prometheus.NewRegistry()
	successProduce := newMetrics(reg, "produce", true)
	failedProduce := newMetrics(reg, "produce", false)
	recordMetrics1s(successProduce)
	recordMetricsRs(failedProduce)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	http.ListenAndServe(":2112", nil)
}
