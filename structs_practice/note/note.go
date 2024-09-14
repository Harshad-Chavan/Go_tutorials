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
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNote(notetitle, notecontent string) (Note, error) {

	if notetitle == "" || notecontent == "" {
		return Note{}, errors.New("title or content can't be empty")
	}

	return Note{
		Title:     notetitle,
		Content:   notecontent,
		CreatedAt: time.Now(),
	}, nil

}

func (note Note) DisplayNote() {
	fmt.Println(note.Title)
	fmt.Println(note.Content)
	fmt.Println(note.CreatedAt)

}

func (note Note) SaveToFile() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename + ".json")

	// marshal will only convert data that is exported i.e starts with capital letter
	json_content, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json_content, 0644)

}
