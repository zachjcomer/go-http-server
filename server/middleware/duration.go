package server

import (
	"log"
	"net/http"
	"time"
)

type DurationMiddleware struct {
	logger *log.Logger
}

func ConfigureDurationMiddleware(logger *log.Logger) *DurationMiddleware {
	if logger == nil {
		panic("No logger!")
	}

	return &DurationMiddleware{logger: logger}
}

func (m *DurationMiddleware) GetHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		m.logger.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}
