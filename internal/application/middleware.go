package application

import (
	"fmt"
	"net/http"
)

// MyLoggingMiddleware adalah middleware untuk logging
func MyLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implementasi logging di sini
		fmt.Println("Logging:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
