package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func NewTodo(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("content can't be empty")
	}

	return Todo{
		Text: content,
	}, nil

}

func (todo Todo) Display() {
	fmt.Println(todo.Text)

}

func (todo Todo) SaveToFile() error {

	filename := "todo" + ".json"

	// marshal will only convert data that is exported i.e starts with capital letter
	json_content, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json_content, 0644)

}
