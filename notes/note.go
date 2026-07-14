package main

import "time"

type note struct {
	Id        int
	Title     string
	Content   string
	CreatedAt time.Time
}

type NoteStore map[string]string

var publicNotes = NoteStore{}


