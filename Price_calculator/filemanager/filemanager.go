package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("Error while opening the file")
	}

	// to read the content of the file
	scanner := bufio.NewScanner(file)

	//
	var lines []string
	//read line by line scanner.scan return bool value..if value is there it will return true
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())

	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Error while reading the file")
	}

	file.Close()
	return lines, nil

}

func WriteJson(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("Error in creating a file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to json")

	}

	file.Close()
	return nil

}
