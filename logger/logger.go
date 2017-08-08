package logger

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func Print(message string, w http.ResponseWriter) {
	start := time.Now()

	fmt.Printf(" %s - %+v\n ", time.Since(start), message)
	if w != nil {
		fmt.Fprint(w, time.Since(start), " - ", message, " ")
	}
}
