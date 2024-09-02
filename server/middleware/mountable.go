package server

import "net/http"

type Mountable interface {
	GetHandler(handler http.Handler) http.Handler
}
