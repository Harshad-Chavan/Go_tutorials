package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
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

func (note Note) SaveToFile() error {
	filename := strings.ReplaceAll(note.title, " ", "_")
	filename = strings.ToLower(filename)

	json_content, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json_content, 0644)

}
