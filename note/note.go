package note

import (
	"fmt"
	"time"
)

type Note struct {
	owner       string
	content     string
	dateCreated time.Time
}

func (n *Note) GetContent() string {
	return n.content
}

type NoteValidationErr struct {
	message string
}

func (e *NoteValidationErr) Error() string {
	return fmt.Sprintf("Cannot construct Note with given arguments. %s", e.message)
}

func NewNote(owner string, content string) (*Note, error) {
	if owner == "" {
		return nil, &NoteValidationErr{"Author must have a name."}
	}

	if content == "" {
		return nil, &NoteValidationErr{"Note must have some content."}
	}

	return &Note{owner, content, time.Now()}, nil
}
