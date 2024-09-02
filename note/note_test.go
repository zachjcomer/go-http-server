package note

import (
	"testing"
)

func TestNewNote(t *testing.T) {
	var owner = "Tester"
	var content = "Hello, world!"
	_, err := NewNote(owner, content)

	if err != nil {
		t.Errorf("newNote")
	}
}

func TestNewNoteNoOwner(t *testing.T) {
	var owner = ""
	var content = "Hello, world!"
	_, err := NewNote(owner, content)

	if err == nil {
		t.Errorf("newNote should fail without an owner.")
	}
}

func TestNewNoteNoContent(t *testing.T) {
	var owner = "Tester"
	var content = ""
	_, err := NewNote(owner, content)

	if err == nil {
		t.Errorf("newNote should fail without content.")
	}
}
