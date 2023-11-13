package prometheus

import (
	"main/internal/common/metrics"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func CollectMetrics(next http.Handler, metrics metrics.Metrics) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.RequestURI
		method := r.Method

		timer := prometheus.NewTimer(metrics.HttpDuration.WithLabelValues(path, method))
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)

		statusCode := rw.statusCode

		metrics.TotalRequests.WithLabelValues(path, method, strconv.Itoa(statusCode)).Inc()

		timer.ObserveDuration()
	})
}
