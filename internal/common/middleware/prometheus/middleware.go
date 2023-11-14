package prometheus

import (
	"main/internal/common/metrics"
	"net/http"
	"strconv"
	"time"
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
		start := time.Now()
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)
		elasped := time.Since(start).Seconds()

		statusCode := strconv.Itoa(rw.statusCode)
		path := r.RequestURI
		method := r.Method

		metrics.TotalRequests.WithLabelValues(path, method, statusCode).Inc()
		metrics.HttpDuration.WithLabelValues(path, method, statusCode).Observe(elasped)
	})
}
