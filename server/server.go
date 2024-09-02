package server

import (
	"fmt"
	// "go-http-server/note"insomin

	"go-http-server/diag"
	middleware "go-http-server/server/middleware"
	"log"
	"net/http"
)

type Server struct {
	handler http.Handler
}

func (s *Server) Mount(middlewares ...middleware.Mountable) {
	for _, m := range middlewares {
		s.handler = m.GetHandler(s.handler)
	}
}

func GetServer(logger *log.Logger) http.Handler {
	mux := http.NewServeMux()

	// mux.Handle("/notes", &note.NotesHandler{})
	mux.HandleFunc("/", getHeartbeat)
	mux.Handle("/diag", diag.LogHandler(logger))

	httpServer := Server{handler: mux}

	// Each middleware wraps the preceeding handler, so the most "bare metal" thing should come first.
	httpServer.Mount(
		middleware.ConfigureDurationMiddleware(logger),
		middleware.ConfigureLoggingMiddleware(logger),
		middleware.ConfigureTracingMiddleware(),
	)

	return httpServer.handler
}

func getHeartbeat(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
