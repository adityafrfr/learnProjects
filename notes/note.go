package main

import (
	"fmt"
	"sync"

	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NoteStore struct {
	mu    sync.RWMutex
	notes map[string]Note
}

var publicNotes = NoteStore{
	notes: make(map[string]Note),
}

func (n *NoteStore) Create(noteObject CreateNoteRequest) Note {
	n.mu.Lock()
	defer n.mu.Unlock()
	note := Note{
		ID:        uuid.NewString(),
		Title:     noteObject.Title,
		Content:   noteObject.Content,
		CreatedAt: time.Now(),
	}
	n.notes[note.ID] = note
	return note
}

func (n *NoteStore) Read() []Note {

	n.mu.RLock()
	defer n.mu.RUnlock()
	var result []Note
	for _, note := range n.notes {
		result = append(result, note)
	}

	return result
}

func (n *NoteStore) ReadById(id string) (Note, error) {

	n.mu.RLock()
	defer n.mu.RUnlock()
	note, ok := n.notes[id]

	if !ok {
		return Note{}, fmt.Errorf("Note with uuid %v not found", id)
	}

	return note, nil
}
