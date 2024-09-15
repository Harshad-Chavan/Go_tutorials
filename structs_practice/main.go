package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

// interface
type saver interface {
	SaveToFile() error
}

// embeded interface
type outputable interface {
	saver
	Display()
}

func getNoteData() (string, string) {
	noteTitle := getUserInput("Enter your note Title")
	noteContent := getUserInput("Enter your note content")

	return noteTitle, noteContent

}

func getTodoData() string {
	todoText := getUserInput("Enter your todo Title")

	return todoText

}

func getUserInput(prompt string) string {
	// problem : the below code will not accept multi word strings
	var text string
	// fmt.Print(prompt)
	// fmt.Scanln(&value)
	// return value

	fmt.Printf("%v : ", prompt)

	// this creates a reader object that reads from the command line
	reader := bufio.NewReader(os.Stdin)

	// this tells the reader object to read till \n or enter
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	// this will remove the \n as it is also included
	text = strings.TrimSuffix(text, "\n")

	// this will remove the \r as it is also included this is for windows as next line in \r\n
	text = strings.TrimSuffix(text, "\r")

	return text
}

func SaveData(data saver) error {
	err := data.SaveToFile()

	if err != nil {

		return err
	}

	return nil
}

func Output(data outputable) error {
	data.Display()
	err := SaveData(data)
	return err

}

func main() {
	// get user input to get the notes from the user
	title, content := getNoteData()
	note, err := note.NewNote(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	note.Display()
	err = SaveData(note)

	if err != nil {
		fmt.Print(err)
		fmt.Print("Saving the note Failed")
		return
	}

	// TODO code
	todoText := getTodoData()
	todo, err := todo.NewTodo(todoText)

	if err != nil {
		fmt.Print(err)
		return
	}

	err = Output(todo)

	if err != nil {
		fmt.Print(err)
		fmt.Print("Saving the todo Failed")
		return
	}

	fmt.Print("Todo saved successfully")

}
