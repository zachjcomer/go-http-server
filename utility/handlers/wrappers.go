package utility

import (
	"fmt"
	utility "go-http-server/utility/serialization"
	"net/http"
)

func Post[T any, U any](f func(t T) U) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request, decodeErr := utility.Decode[T](w, r)
		if decodeErr != nil {
			BadRequest(w, r)
			return
		}

		response := f(request)

		if err := utility.Encode[U](&response, w); err != nil {
			InternalServerError(w, r)
			return
		}
	})
}

func Get[U any](f func() U) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := f()

		if err := utility.Encode[U](&response, w); err != nil {
			InternalServerError(w, r)
			return
		}
	})
}

func Pipe(f any) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := utility.PipeHttpToFunction(f, w, r); err != nil {
			fmt.Println(err)

			// how to know which? standardized errors?
			InternalServerError(w, r)
			// BadRequest(w, r)

			return
		}

		Ok(w, r)
	})
}
