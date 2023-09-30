package common_middleware

import (
	"io"
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		log.Println(req.Header)

		next.ServeHTTP(w, req)

		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
		log.Println(w.Header())
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
