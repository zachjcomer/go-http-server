package utility

import "net/http"

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func Ok(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
