package metrics

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Middleware struct {
	handlers handlers
	metrics  metrics
}

func NewMiddleware(h handlers, m metrics) Middleware {
	return Middleware{
		handlers: h,
		metrics:  m,
	}
}

func (m *Middleware) preprocessingPath(path string) string {
	trimmedPath := strings.TrimSuffix(strings.TrimPrefix(path, "/api/v1"), "/")
	re := regexp.MustCompile(`(\d+)`)
	return re.ReplaceAllString(trimmedPath, "{id}")
}

func (m *Middleware) Collecting(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/wave" {
			start := time.Now()
			rw := NewResponseWriter(w)
			next.ServeHTTP(rw, r)
			elasped := time.Since(start).Seconds()

			statusCode := strconv.Itoa(rw.statusCode)
			path := m.preprocessingPath(r.RequestURI)
			handler := m.handlers.getHandler(path)
			method := r.Method

			m.metrics.TotalRequests.WithLabelValues(path, method, statusCode, handler).Inc()
			m.metrics.HttpDuration.WithLabelValues(path, method, statusCode, handler).Observe(elasped)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
