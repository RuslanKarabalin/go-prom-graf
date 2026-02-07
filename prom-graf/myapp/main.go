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

func newPromCounter(reg prometheus.Registerer, name string, event string, isSuccess bool) *metrics {
	reg = prometheus.WrapRegistererWith(prometheus.Labels{"event": event, "isSuccess": strconv.FormatBool(isSuccess)}, reg)
	m := &metrics{
		opsProcessed: promauto.With(reg).NewCounter(prometheus.CounterOpts{
			Name: name,
		}),
	}
	return m
}

func newPromGauge(reg prometheus.Registerer, name string) *metrics {
	m := &metrics{
		opsProcessed: promauto.With(reg).NewGauge(prometheus.GaugeOpts{
			Name: name,
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
			m.opsProcessed.Add(2 * float64(r))
			time.Sleep(time.Duration(r) * time.Second)
		}
	}()
}

func main() {
	reg := prometheus.NewRegistry()
	successProduce := newPromCounter(reg, "myapp_counter_total", "produce", true)
	failedProduce := newPromCounter(reg, "myapp_counter_total", "produce", false)

	successConsume := newPromCounter(reg, "myapp_counter_total", "consume", true)
	failedConsume := newPromCounter(reg, "myapp_counter_total", "consume", false)

	gauge := newPromGauge(reg, "myapp_gauge_total")

	recordMetrics1s(successProduce)
	recordMetricsRs(failedProduce)

	recordMetricsRs(successConsume)
	recordMetrics2s(failedConsume)

	recordMetricsRs(gauge)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	http.ListenAndServe(":2112", nil)
}
