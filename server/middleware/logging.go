package server

import (
	"log"
	"net/http"
)

type LoggingMiddleware struct {
	logger *log.Logger
}

func ConfigureLoggingMiddleware(logger *log.Logger) *LoggingMiddleware {
	if logger == nil {
		panic("No logger!")
	}

	return &LoggingMiddleware{logger: logger}
}

func (m *LoggingMiddleware) GetHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.logger.Printf("%s", r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
