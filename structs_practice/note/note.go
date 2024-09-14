package note

import (
	"errors"
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

func DisplayNote() {

}
