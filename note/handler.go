package note

import (
	"net/http"
	"regexp"
)

var (
	NotesUri       = regexp.MustCompile(`^/notes/*$`)
	NotesWithIdUri = regexp.MustCompile(`^/notes/`)
)

type NotesHandler struct{}

// func (h *NotesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch {
// 	case r.Method == http.MethodGet && NotesUri.MatchString(r.URL.Path):
// 		handler := server.Get(func() Note {
// 			note, err := NewNote("Test", "Hello, world!")
// 			if err != nil {

// 			}

// 			return note
// 		})

// 		handler.ServeHTTP(w, r)
// 	}
// }

func (h *NotesHandler) getNotes(w http.ResponseWriter, r *http.Request) {
	// dummyNote, err := NewNote("Test", "Helloh")
	// if (err != nil)
}
