package main

import (
	"errors"
	"fmt"
)

func getNoteData() (string, string, error) {
	noteTitle, err := getUserInput("Enter your note Title")

	if err != nil {
		return "", "", err
	}

	noteContent, err := getUserInput("Enter your note content")

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return noteTitle, noteContent, nil

}

func getUserInput(prompt string) (string, error) {
	var value string
	fmt.Print(prompt)
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("Value can't be empty.")
	}
	return value, nil
}

func main() {
	// get user input to get the notes from the user
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
	}

}
