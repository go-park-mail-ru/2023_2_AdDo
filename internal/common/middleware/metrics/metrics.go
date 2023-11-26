package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type metrics struct {
	TotalRequests  *prometheus.CounterVec
	HttpDuration   *prometheus.HistogramVec
}

func NewMetrics(reg prometheus.Registerer) metrics {
	const namePrefix = "musicon_"
	m := metrics{
		TotalRequests: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: namePrefix + "requests_total",
				Help: "Number of get requests.",
			},
			[]string{"path", "method", "code", "handler"},
		),
		HttpDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: namePrefix + "response_time_seconds",
				Help: "Duration of HTTP requests.",
			},
			[]string{"path", "method", "code", "handler"},
		),
	}
	reg.MustRegister(m.TotalRequests, m.HttpDuration)
	return m
}
