package utility

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Encode[T any](response *T, w http.ResponseWriter) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return fmt.Errorf("error encoding response: %w", err)
	}

	return nil
}

func Decode[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var t T
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		return t, fmt.Errorf("error decoding request: %w", err)
	}

	return t, nil
}
