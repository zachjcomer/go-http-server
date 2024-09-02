package utility

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ideally split this into two layers, one that knows about http and one that just knows about readers/writers

// Attempts to pipe the HTTP request into named, typed parameters declared by f, which is
// part of a http.Handler.
//
// The body, if present, is the only part of the request that may serve a struct.
// Thus the intent is that f forms a closure over any other non-request structs via closure.
func PipeHttpToFunction(f interface{}, w http.ResponseWriter, r *http.Request) error {
	fType := reflect.TypeOf(f)
	if fType.Kind() != reflect.Func {
		return fmt.Errorf("unable to parse handler into func")
	}

	input := make([]reflect.Value, fType.NumIn())

	for i := 0; i < fType.NumIn(); i++ {
		param := fType.In(i)

		// struct is only matched with body, body only matched with struct
		if param.Kind() == reflect.Struct {
			p, err := decode(w, r, param)
			if err != nil {
				return fmt.Errorf("error parsing body: %w", err)
			}

			input[i] = p
		} else if param.Kind() == reflect.Pointer {
			p, err := decode(w, r, param.Elem())
			if err != nil {
				return fmt.Errorf("error parsing body: %w", err)
			}

			input[i] = p.Addr()
		}
		// } else if isKeyInHeaders(r, param) {
		// 	paramKey := formatHeader(paramN)

		// 	val := r.Header[paramKey]

		// 	input[i] = reflect.ValueOf(val)
		// }
	}

	fVal := reflect.ValueOf(f)
	fVal.Call(input)

	// write response

	return nil
}

// name match + value castable to key type and/or array match of castable type
func isKeyInHeaders(r *http.Request, key reflect.Type) bool {
	paramKey := formatHeader(key.Name())

	if r.Header[paramKey] == nil {
		return false
	}

	if len(r.Header[paramKey]) > 1 && key.Kind() != reflect.Array {
		return false
	}

	// still need to handle type assertions of key

	return true
}

func decode(w http.ResponseWriter, r *http.Request, t reflect.Type) (reflect.Value, error) {
	targ := reflect.New(t).Elem()

	if err := json.NewDecoder(r.Body).Decode(targ.Addr().Interface()); err != nil {
		return reflect.Zero(t), fmt.Errorf("error decoding request :%w", err)
	}

	return targ, nil
}

func formatHeader(key string) string {
	return cases.Title(language.Und).String(key)
}

// match body -> should be only struct param in handler
// fill in remainder with headers and shit
