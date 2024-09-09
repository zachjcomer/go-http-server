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

	if fType.NumIn() != 1 || fType.In(0).Kind() != reflect.Struct {
		return fmt.Errorf("func must have a single struct parameter to associate request with fields")
	}

	req := reflect.New(fType.In(0))     // Create a struct that will be passed to the func
	ref := reflect.ValueOf(&req).Elem() // A reference that can be modified

	fmt.Println(req.Type().Elem())

	for i := 0; i < req.Type().Elem().NumField(); i++ {
		field := req.Type().Elem().Field(i)
		fmt.Println(field.Name, field.Type.Name())
	}

	for i := 0; i < ref.NumField(); i++ {
		t := ref.Type().Field(i) // Type info to get field name
		field := ref.Field(i)    // Field to set

		fmt.Println("Can set ", t.Name, ": ", field.CanSet())

		if field.Kind() == reflect.Struct {
			p, err := decode(w, r, field.Type())
			if err != nil {
				return fmt.Errorf("error parsing body: %w", err)
			}

			field.Set(p)
		} else if field.Kind() == reflect.Pointer {
			p, err := decode(w, r, field.Type().Elem())
			if err != nil {
				return fmt.Errorf("error parsing body: %w", err)
			}

			field.Set(p)
		} else if isFieldInHeaders(r, t) {
			field.Set(reflect.ValueOf(r.Header[t.Name]))
		}
	}

	fVal := reflect.ValueOf(f)
	fVal.Call([]reflect.Value{req})

	// write response

	return nil
}

// name match + value castable to key type and/or array match of castable type
func isFieldInHeaders(r *http.Request, key reflect.StructField) bool {
	paramKey := formatHeader(key.Name)

	if r.Header[paramKey] == nil {
		return false
	}

	if len(r.Header[paramKey]) > 1 && key.Type.Kind() != reflect.Array {
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
