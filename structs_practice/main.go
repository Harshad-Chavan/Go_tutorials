package main

import (
	"fmt"

	"example.com/notes/note"
)

func getNoteData() (string, string) {
	noteTitle := getUserInput("Enter your note Title")
	noteContent := getUserInput("Enter your note content")

	return noteTitle, noteContent

}

func getUserInput(prompt string) string {
	var value string
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
}

func main() {
	// get user input to get the notes from the user
	title, content := getNoteData()
	note, err := note.NewNote(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(note)

}
