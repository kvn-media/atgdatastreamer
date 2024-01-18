package application

import (
	"log"
	"net/http"
	"time"
)

// MyLoggingMiddleware is middleware for logging
func MyLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log information about the request
		statusCode := w.(interface {
			Status() int
		}).Status()
		elapsed := time.Since(startTime)

		log.Printf(
			"Method: %s\tPath: %s\tStatus: %d\tElapsed: %s\n",
			r.Method, r.URL.Path, statusCode, elapsed,
		)
	})
}