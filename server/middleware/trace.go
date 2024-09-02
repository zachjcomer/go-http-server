package server

import (
	"net/http"

	"github.com/google/uuid"
)

type TracingMiddleware struct{}

func ConfigureTracingMiddleware() *TracingMiddleware {
	return &TracingMiddleware{}
}

func (m *TracingMiddleware) GetHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Trace-Id", uuid.NewString())
		handler.ServeHTTP(w, r)
	})
}
