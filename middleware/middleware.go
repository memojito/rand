package middleware

import (
	"log"
	"net/http"
	"time"
)

type loggingMiddleware struct {
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf(r.Method, r.URL.Path, time.Since(start))
	})
}
