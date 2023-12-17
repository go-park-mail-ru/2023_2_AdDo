package common

import (
	"github.com/sirupsen/logrus"
	"io"
	"main/internal/common/utils"
	"net/http"
	"time"
)

func Logging(next http.Handler, logger *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		logger.WithFields(logrus.Fields{
			"request_id":     utils.GenReqId(req.RequestURI + req.Method),
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

func PanicRecovery(next http.Handler, logger *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		defer func(w http.ResponseWriter, request *http.Request) {
			if err := recover(); err != nil {
				logger.WithFields(logrus.Fields{
					"request id":     utils.GenReqId(request.RequestURI + request.Method),
					"request_method": request.Method,
					"request_header": request.Header,
					"request_uri":    request.RequestURI,
					"err":            err,
				}).Infoln("panic happened")

				w.WriteHeader(http.StatusInternalServerError)
				_, err := io.WriteString(w, `{"status": 500, "err": "Unknown error"}`)
				if err != nil {
					logger.Errorln(err.Error())
				}
			}
		}(w, request)

		next.ServeHTTP(w, request)
	})
}
