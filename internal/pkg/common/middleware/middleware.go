package common_middleware

import (
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

func Logging(next http.Handler, logger *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		logger.WithFields(logrus.Fields{
			"request_method": req.Method,
			"request_uri":    req.RequestURI,
			"request header": req.Header,
			"request_body":   req.Body,
		}).Infoln("start request processing")

		next.ServeHTTP(w, req)

		logger.WithFields(logrus.Fields{
			"time_since_start": time.Since(start),
			"response_header":  w.Header(),
		}).Infoln("request processed")
	})
}

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		defer func(w http.ResponseWriter, request *http.Request) {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				io.WriteString(w, `{"status": 500, "err": "Unknown error"}`)
			}
		}(w, request)
		next.ServeHTTP(w, request)
	})
}
