package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, requrest *http.Request) {
		start := time.Now()

		next.ServeHTTP(writer, requrest)

		log.Printf("%s %s %s", requrest.Method, requrest.URL.Path, time.Since(start))
	})
}
