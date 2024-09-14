package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func NewNote(notetitle, notecontent string) (Note, error) {

	if notetitle == "" || notecontent == "" {
		return Note{}, errors.New("title or content can't be empty")
	}

	return Note{
		title:     notetitle,
		content:   notecontent,
		createdAt: time.Now(),
	}, nil

}

func (note Note) DisplayNote() {
	fmt.Println(note.title)
	fmt.Println(note.content)
	fmt.Println(note.createdAt)

}
