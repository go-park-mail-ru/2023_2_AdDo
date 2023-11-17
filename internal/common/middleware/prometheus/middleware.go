package prometheus

import (
	"main/internal/common/metrics"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func preprocessingPath(path string) string {
	trimmedPath := strings.TrimSuffix(strings.TrimPrefix(path, "/api/v1"), "/")
	re := regexp.MustCompile(`(\d+)`)
	return re.ReplaceAllString(trimmedPath, "{id}")
}

func CollectMetrics(next http.Handler, metrics metrics.Metrics, handlers handlersMap) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)
		elasped := time.Since(start).Seconds()

		statusCode := strconv.Itoa(rw.statusCode)
		path := preprocessingPath(r.RequestURI)
		handler := handlers.getHandler(path)
		method := r.Method

		metrics.TotalRequests.WithLabelValues(path, method, statusCode, handler).Inc()
		metrics.HttpDuration.WithLabelValues(path, method, statusCode, handler).Observe(elasped)
	})
}

