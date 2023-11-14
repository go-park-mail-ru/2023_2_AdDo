package prometheus

import (
	"main/internal/common/metrics"
	"net/http"
	"regexp"
	"strconv"
	"strings"
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



// TODO get handler from request context
func CollectMetrics(next http.Handler, metrics metrics.Metrics, handlers map[string]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)
		elasped := time.Since(start).Seconds()

		statusCode := strconv.Itoa(rw.statusCode)
		path := strings.TrimSuffix(strings.TrimPrefix(r.RequestURI, "/api/v1"), "/")
		handler := getHandler(handlers, path)
		method := r.Method

		metrics.TotalRequests.WithLabelValues(path, method, statusCode, handler).Inc()
		metrics.HttpDuration.WithLabelValues(path, method, statusCode, handler).Observe(elasped)
	})
}

const (
	userHandler     = "user"
	trackHandler    = "track"
	albumHandler    = "album"
	artistHandler   = "artist"
	playlistHandler = "playlist"
)

// cringe
func getHandler(handlers map[string]string, path string) string {
	if handler, ok := handlers[path]; ok {
		return handler
	}

	templates := map[string]string{
		"/album.*":    albumHandler,
		"/artist.*":   artistHandler,
		"/playlist.*": playlistHandler,
		"/track.*":    trackHandler,
	}
	for pattern, handler := range templates {
		if res, _ := regexp.MatchString(pattern, path); res {
			return handler
		}
	}
	return ""
}

func NewHandlersMap() map[string]string {
	handlers := make(map[string]string)
	userPathTemplates := []string{
		"/sign_up",
		"/login",
		"/update_info",
		"/upload_avatar",
		"/remove_avatar",
		"/auth",
		"/me",
		"/logout",
	}
	trackPathTemplates := []string{
		"/listen",
		"/collection/tracks",
	}
	albumPathTempaltes := []string{
		"/feed",
		"/new",
		"/most_liked",
		"/popular",
	}

	for _, template := range userPathTemplates {
		handlers[template] = userHandler
	}
	for _, template := range trackPathTemplates {
		handlers[template] = trackHandler
	}
	for _, template := range albumPathTempaltes {
		handlers[template] = albumHandler
	}
	return handlers
}
